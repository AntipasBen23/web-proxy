package main

import (
	"log"
	"net/http"

	"github.com/AntipasBen23/web-proxy/services/auth-service/internal/config"
	"github.com/AntipasBen23/web-proxy/services/auth-service/internal/middleware"
)

func main() {
	cfg := config.Load()

	mux := http.NewServeMux()

	mux.HandleFunc("/validate", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Authenticated"))
	})

	handler := middleware.AuthMiddleware(cfg, mux)

	log.Printf("🔐 Auth service running on :%s", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, handler))
}
