package auth

import "github.com/golang-jwt/jwt/v5"

type IDTokenClaims struct {
	Iss           string `json:"iss"`
	Azp           string `json:"azp"`
	Aud           string `json:"aud"`
	Sub           string `json:"sub"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	AtHash        string `json:"at_hash"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Iat           int    `json:"iat"`
	Exp           int    `json:"exp"`
	jwt.RegisteredClaims
}

type LoginUrlWithState struct {
	Url   string
	State string
}
