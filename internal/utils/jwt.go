package utils

import (
	"errors"
	"log"

	"github.com/golang-jwt/jwt/v5"
	"github.com/ilyasa1211/go-google-openid/internal/config/auth"
	"github.com/ilyasa1211/go-google-openid/internal/core/domain/user"
)

type UserClaim struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
	jwt.RegisteredClaims
}

var config auth.JWTConfig = *auth.NewJWTConfig()

func GenJWTToken(u *user.User) string {
	s, err := jwt.NewWithClaims(jwt.SigningMethodHS512, &UserClaim{
		ID:   u.ID,
		Name: u.Name,
	}).SignedString([]byte(config.Secret))

	if err != nil {
		log.Fatalln("error signing jwt", err)
	}

	return s
}

func VerifyJWTToken(tokenStr string) (*UserClaim, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &UserClaim{}, func(t *jwt.Token) (any, error) {
		return []byte(config.Secret), nil
	}, jwt.WithValidMethods([]string{
		jwt.SigningMethodHS512.Alg(),
	}))

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*UserClaim); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token claims")
}
