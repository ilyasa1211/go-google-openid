package middlewares

import "net/http"

func ChainMiddlewares(handler http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}

	return handler
}
