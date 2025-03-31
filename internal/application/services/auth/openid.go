package auth

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/ilyasa1211/go-google-openid/internal/application/dto"
	authInterface "github.com/ilyasa1211/go-google-openid/internal/core/domain/auth"
	"github.com/ilyasa1211/go-google-openid/internal/core/domain/user"
	"github.com/ilyasa1211/go-google-openid/internal/utils"
)

type OpenIdAuthService struct {
	AuthService
	Strategy authInterface.OpenIdStrategy
	Cache    authInterface.AuthCache
}

func NewOpenIdAuthService(r user.UserRepository, s authInterface.OpenIdStrategy, c authInterface.AuthCache) *OpenIdAuthService {
	return &OpenIdAuthService{AuthService: AuthService{r}, Strategy: s, Cache: c}
}

func (s *OpenIdAuthService) GetLoginUrl() (string, error) {
	res, err := s.Strategy.GetLoginUrl()

	if err != nil {
		return "", err
	}

	key := strings.Join([]string{
		"openid",
		res.State,
	}, "-")

	s.Cache.Set(key, "ok")

	return res.Url, nil
}

func (s *OpenIdAuthService) HandleLoginCallback(r *http.Request) (string, error) {
	state := r.URL.Query().Get("state")

	key := strings.Join([]string{
		"openid",
		state,
	}, "-")

	if s.Cache.Get(key) == "" {
		return "", fmt.Errorf("state does not match: ")
	}

	// s.Cache.Del(key)

	claims := s.Strategy.HandleLoginCallback(r)

	createUser := &dto.CreateUserRequest{
		Name:     claims.Name,
		Email:    claims.Email,
		Password: "",
	}

	if err := s.AuthService.r.Create(createUser); err != nil {
		return "", fmt.Errorf("failed to create user %w", err)
	}

	token := utils.GenJWTToken(&user.User{
		Name:  createUser.Name,
		Email: createUser.Email,
	})

	return token, nil
}
