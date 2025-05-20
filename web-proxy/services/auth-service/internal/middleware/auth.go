package middleware

import (
	"net/http"
	"strings"

	"github.com/AntipasBen23/web-proxy/services/auth-service/internal/config"
	"github.com/AntipasBen23/web-proxy/services/auth-service/internal/validator"
)

func AuthMiddleware(cfg *config.Config, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get(cfg.AuthHeaderName)
		if token == "" {
			http.Error(w, "Missing auth header", http.StatusUnauthorized)
			return
		}

		if !validator.Validate(validator.Strategy(cfg.AuthStrategy), token, cfg.AuthSecretKey) {
			http.Error(w, "Invalid credentials", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
