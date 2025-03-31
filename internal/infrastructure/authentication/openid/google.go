package openid

import (
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/big"
	"net/http"
	"net/url"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	config "github.com/ilyasa1211/go-google-openid/internal/config/auth/openid"
	"github.com/ilyasa1211/go-google-openid/internal/core/domain/auth"
	"github.com/ilyasa1211/go-google-openid/internal/utils"
)

type DiscoveryDocument struct {
	Issuer                            string   `json:"issuer"`
	AuthorizationEndpoint             string   `json:"authorization_endpoint"`
	DeviceAuthorizationEndpoint       string   `json:"device_authorization_endpoint"`
	TokenEndpoint                     string   `json:"token_endpoint"`
	UserinfoEndpoint                  string   `json:"userinfo_endpoint"`
	RevocationEndpoint                string   `json:"revocation_endpoint"`
	JwksURI                           string   `json:"jwks_uri"`
	ResponseTypesSupported            []string `json:"response_types_supported"`
	SubjectTypesSupported             []string `json:"subject_types_supported"`
	IDTokenSigningAlgValuesSupported  []string `json:"id_token_signing_alg_values_supported"`
	ScopesSupported                   []string `json:"scopes_supported"`
	TokenEndpointAuthMethodsSupported []string `json:"token_endpoint_auth_methods_supported"`
	ClaimsSupported                   []string `json:"claims_supported"`
	CodeChallengeMethodsSupported     []string `json:"code_challenge_methods_supported"`
	GrantTypesSupported               []string `json:"grant_types_supported"`
}

type SuccessResponseExchangeToken struct {
	AccessToken           string `json:"access_token"`
	IDToken               string `json:"id_token"`
	ExpiresIn             int    `json:"expires_in"`
	RefreshTokenExpiresIn int    `json:"refresh_token_expires_in"`
	TokenType             string `json:"token_type"`
	Scope                 string `json:"scope"`
	RefreshToken          string `json:"refresh_token"`
}

type GoogleOpenIdAuthentication struct {
}

type JWK struct {
	Kty string `json:"kty"`
	Kid string `json:"kid"`
	E   string `json:"e"`
	N   string `json:"n"`
	Alg string `json:"alg"`
	Use string `json:"use"`
}

type JWKS struct {
	Keys []JWK `json:"keys"`
}

type UserClaims struct {
	jwt.RegisteredClaims
}

var conf *config.GoogleOpenIdConf = config.NewGoogleOpenIdConfg()
var docs *DiscoveryDocument = parseDiscoveryDocument(conf.DiscoveryDocumentUrl)

func NewGoogleOpenIdAuthentication() *GoogleOpenIdAuthentication {
	return &GoogleOpenIdAuthentication{}
}

func (a *GoogleOpenIdAuthentication) GetLoginUrl() (*auth.LoginUrlWithState, error) {
	parsedUrl, err := url.Parse(docs.AuthorizationEndpoint)

	if err != nil {
		return nil, fmt.Errorf("error parsing url: %w", err)
	}

	state := utils.GenerateRandomState()

	q := parsedUrl.Query()
	q.Set("state", state)
	q.Set("client_id", conf.ClientId)
	q.Set("prompt", "select_account")
	q.Set("response_type", "code")
	// q.Set("nonce")
	q.Set("redirect_uri", conf.CallbackUrl)
	q.Set("scope", strings.Join([]string{
		"https://www.googleapis.com/auth/userinfo.email",
		"https://www.googleapis.com/auth/userinfo.profile",
		"openid",
	}, " "))

	parsedUrl.RawQuery = q.Encode()

	return &auth.LoginUrlWithState{
		Url:   parsedUrl.String(),
		State: state,
	}, nil
}

func GetJWK(j *JWKS, kid string) *JWK {
	for _, jw := range j.Keys {
		if jw.Kid == kid {
			return &jw
		}
	}

	return nil
}
func (j *JWK) ToPublicKey() *rsa.PublicKey {
	decodedE, _ := base64.RawURLEncoding.DecodeString(j.E)
	decodedN, _ := base64.RawURLEncoding.DecodeString(j.N)

	n := new(big.Int).SetBytes(decodedN)
	e := new(big.Int).SetBytes(decodedE)

	return &rsa.PublicKey{
		E: int(e.Int64()),
		N: n,
	}

}
func (a *GoogleOpenIdAuthentication) HandleLoginCallback(r *http.Request) *auth.IDTokenClaims {
	q := r.URL.Query()

	code := q.Get("code")

	resp := exchangeCodeForToken(code)

	// get id token and validate
	jwks := fetchJWKS()
	// token, err := jwt.Parse(resp.IDToken, nil)
	token, _, err := jwt.NewParser().ParseUnverified(resp.IDToken, &auth.IDTokenClaims{})

	if err != nil {
		log.Fatalln("Invalid jwt")
	}

	_, ok := token.Method.(*jwt.SigningMethodRSA)

	if !ok {
		log.Fatalln("Invalid signing method")
	}

	kid, _ := token.Header["kid"].(string)

	jwk := GetJWK(jwks, kid)
	rsaPubKey := jwk.ToPublicKey()

	if err := verifyToken(resp.IDToken, rsaPubKey); err != nil {
		log.Fatalln("token verify failed: ", err)
	}

	userClaims, _ := token.Claims.(*auth.IDTokenClaims)

	return userClaims
}

func verifyToken(jwtToken string, pubkey *rsa.PublicKey) error {
	_, err := jwt.Parse(jwtToken, func(t *jwt.Token) (interface{}, error) {
		return pubkey, nil
	}, jwt.WithValidMethods([]string{
		jwt.SigningMethodRS256.Alg(),
	}),
		jwt.WithIssuer(docs.Issuer),
	)

	if err != nil {
		return fmt.Errorf("error validating token: %w", err)
	}

	return nil
}

func fetchJWKS() *JWKS {
	resp, err := http.Get(docs.JwksURI)

	if err != nil {
		log.Fatalln("failed to fetch jwks: ", err)
	}

	defer resp.Body.Close()

	var jwks JWKS

	json.NewDecoder(resp.Body).Decode(&jwks)

	return &jwks
}

func exchangeCodeForToken(code string) *SuccessResponseExchangeToken {
	resp, err := http.PostForm(docs.TokenEndpoint, url.Values{
		"code":          {code},
		"client_id":     {conf.ClientId},
		"client_secret": {conf.ClientSecret},
		"redirect_uri":  {conf.CallbackUrl},
		"grant_type":    {"authorization_code"},
	})

	if err != nil {
		log.Fatalln("error exchange code for tokens: ", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)

		if err != nil {
			log.Fatalln("error exchange code for tokens: ", err)
		}

		log.Fatalln("error exchange code for token", string(body))
	}

	// if resp.StatusCode != http.StatusOK {
	// 	log.Fatalln("error exchange code for tokens: ", err)
	// }

	var response SuccessResponseExchangeToken

	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(&response)

	return &response
}

func parseDiscoveryDocument(url string) *DiscoveryDocument {
	var result DiscoveryDocument

	resp, err := http.Get(url)

	if err != nil {
		log.Fatalln("error fetching discovery document: ", err)
	}

	json.NewDecoder(resp.Body).Decode(&result)

	return &result
}
