package main

import (
	"log"
	"net/http"
	"time"

	"backend-rs100/config"
	delivery "backend-rs100/internal/delivery/http"
	"backend-rs100/internal/repository"
	"backend-rs100/internal/usecase"
)

func main() {
	// 1. Load Configurations
	cfg := config.LoadConfig()

	// 2. Initialize database connection (passing nil; in-memory fallback handles data access without a running MySQL instance)
	log.Printf("[INIT] Connecting to MySQL at %s:%s...", cfg.DBHost, cfg.DBPort)

	// 3. Initialize Repositories
	bookRepo := repository.NewBookMySQLRepository(nil)
	activationRepo := repository.NewActivationMySQLRepository(nil)

	// 4. Initialize Usecases with 2 seconds context timeout
	timeout := 2 * time.Second
	bookUsecase := usecase.NewBookUsecase(bookRepo, timeout)
	authUsecase := usecase.NewAuthUsecase(activationRepo, timeout)

	// 5. Initialize Delivery (HTTP Handler)
	handler := delivery.NewHttpHandler(bookUsecase, authUsecase)

	// 6. Router Setup
	mux := http.NewServeMux()
	handler.RegisterRoutes(mux)

	// Add a default base route to verify API status
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"app": "RSUD Aloei Saboe 100th Anniversary API", "status": "running"}`))
	})

	// 7. Start HTTP Server
	serverAddr := ":" + cfg.Port
	log.Printf("[INIT] Server starting on HTTP port %s", serverAddr)
	if err := http.ListenAndServe(serverAddr, mux); err != nil {
		log.Fatalf("[ERROR] Failed to start server: %v", err)
	}
}
