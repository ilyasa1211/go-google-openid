package middlewares

import (
	"net/http"
	"strings"

	"github.com/ilyasa1211/go-google-openid/internal/utils"
)

func AuthMiddleware() Middleware {
	return func(hf http.HandlerFunc) http.HandlerFunc {

		return func(w http.ResponseWriter, r *http.Request) {

			authHeader := r.Header.Get("Authorization")

			if authHeader == "" {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}

			if !strings.HasPrefix(authHeader, "Bearer") {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}

			auth := strings.Split(authHeader, " ")

			if len(auth) < 2 {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}

			token := auth[1]

			// claims, err
			_, err := utils.VerifyJWTToken(token)

			if err != nil {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}

			hf(w, r)
		}
	}
}
