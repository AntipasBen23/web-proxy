package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
	"github.com/AntipasBen23/web-proxy/services/auth-service/internal/validator"
)

var (
	authHeaderName string
	authSecretKey  string
	authStrategy   string
)

type validationResponse struct {
	Valid   bool   `json:"valid"`
	Message string `json:"message"`
}

func loadEnv() {
	authHeaderName = getEnv("AUTH_HEADER_NAME", "Authorization")
	authSecretKey = getEnv("AUTH_SECRET_KEY", "")
	authStrategy = getEnv("AUTH_STRATEGY", "api_key")

	if authSecretKey == "" {
		log.Fatal("❌ AUTH_SECRET_KEY is required but not set")
	}
	log.Println("✅ Environment loaded")
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func validateHandler(w http.ResponseWriter, r *http.Request) {
    cfg := config.Load()
    token := r.Header.Get(cfg.AuthHeaderName)

    if token == "" || !validator.Validate(validator.Strategy(cfg.AuthStrategy), token, cfg.AuthSecretKey) {
        w.WriteHeader(http.StatusUnauthorized)
        w.Write([]byte("Unauthorized"))
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Authenticated"))
}
	switch authStrategy {
	case "api_key":
		if token != authSecretKey {
			writeJSON(w, http.StatusForbidden, validationResponse{
				Valid:   false,
				Message: "Invalid API key",
			})
			return
		}
	case "token":
		token = strings.TrimPrefix(token, "Bearer ")
		if token != authSecretKey {
			writeJSON(w, http.StatusForbidden, validationResponse{
				Valid:   false,
				Message: "Invalid token",
			})
			return
		}
	default:
		writeJSON(w, http.StatusInternalServerError, validationResponse{
			Valid:   false,
			Message: "Unsupported auth strategy",
		})
		return
	}

	writeJSON(w, http.StatusOK, validationResponse{
		Valid:   true,
		Message: "Authenticated",
	})
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

func main() {
	loadEnv()

	http.HandleFunc("/validate", validateHandler)

	port := getEnv("PORT", "8081")
	log.Printf("🔐 Auth service running on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
