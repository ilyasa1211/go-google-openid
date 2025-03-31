package auth

import "os"

type JWTConfig struct {
	Secret string
}

func NewJWTConfig() *JWTConfig {
	return &JWTConfig{
		Secret: os.Getenv("JWT_SECRET"),
	}
}
