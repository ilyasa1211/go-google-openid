package main

import (
	"fmt"
	"net/http"

	"github.com/ilyasa1211/go-google-openid/internal/application/services/auth"
	"github.com/ilyasa1211/go-google-openid/internal/infrastructure/authentication/openid"
	"github.com/ilyasa1211/go-google-openid/internal/infrastructure/cache/valkey"
	"github.com/ilyasa1211/go-google-openid/internal/infrastructure/database/pgsql"
	"github.com/ilyasa1211/go-google-openid/internal/infrastructure/database/pgsql/repositories"
	auth_handler "github.com/ilyasa1211/go-google-openid/internal/infrastructure/http/handlers/auth"
)

func main() {
	const PORT = 8080

	db := pgsql.NewPGSQLConn()
	cacheClient := valkey.NewValkeyClient()

	cache := valkey.NewCacheImpl(cacheClient)
	userRepository := repositories.NewUserRepository(db)
	authOpenIdStrategy := openid.NewGoogleOpenIdAuthentication()
	authOpenIdService := auth.NewOpenIdAuthService(userRepository, authOpenIdStrategy, cache)
	authOpenIdHandler := auth_handler.NewOpenIDHandler(authOpenIdService)

	http.HandleFunc("GET /login/openid/google", authOpenIdHandler.Login)
	http.HandleFunc("GET /login/openid/google/callback", authOpenIdHandler.HandleLoginCallback)

	http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil)
}
