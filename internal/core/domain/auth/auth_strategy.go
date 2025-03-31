package auth

import "net/http"

type OpenIdStrategy interface {
	GetLoginUrl() (*LoginUrlWithState, error)
	HandleLoginCallback(r *http.Request) *IDTokenClaims
}
