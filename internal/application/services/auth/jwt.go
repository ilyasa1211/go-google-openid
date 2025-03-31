package auth

import (
	"encoding/json"
	"net/http"

	"github.com/ilyasa1211/go-google-openid/internal/application/dto"
	"github.com/ilyasa1211/go-google-openid/internal/core/domain/user"
	"github.com/ilyasa1211/go-google-openid/internal/utils"
)

type JwtAuthService struct {
	AuthService
}

func NewJwtAuthService(r user.UserRepository) *JwtAuthService {
	return &JwtAuthService{AuthService: AuthService{r}}
}

func (s *JwtAuthService) Login(r *http.Request) (string, error) {
	var dto dto.LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		return "", err
	}

	user := s.r.FindByEmail(dto.Email)

	if err := utils.VerifyPassword(dto.Password, user.Password); err != nil {
		return "", err
	}

	return utils.GenJWTToken(user), nil
}

func (s *JwtAuthService) Register(r *http.Request) (string, error) {
	var reg dto.RegisterRequest
	var userCreate dto.CreateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&reg); err != nil {
		return "", err
	}

	userCreate = reg.CreateUserRequest

	userCreate.Password = utils.HashPass(reg.Password)

	if err := s.r.Create(&userCreate); err != nil {
		return "", err
	}

	user := s.r.FindByEmail(reg.Email)

	if err := utils.VerifyPassword(reg.Password, user.Password); err != nil {
		return "", err
	}

	return utils.GenJWTToken(user), nil
}
