package auth

import "net/http"

type JwtService interface {
	Login(r *http.Request) (string, error)
	Register(r *http.Request) (string, error)
}

type OpenIdService interface {
	GetLoginUrl() (string, error)
	HandleLoginCallback(r *http.Request) (string, error)
}
