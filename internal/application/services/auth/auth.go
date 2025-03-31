package auth

import "github.com/ilyasa1211/go-google-openid/internal/core/domain/user"

type AuthService struct {
	r user.UserRepository
}
