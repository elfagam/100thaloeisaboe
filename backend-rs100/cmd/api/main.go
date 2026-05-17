package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"backend-rs100/config"
	delivery "backend-rs100/internal/delivery/http"
	"backend-rs100/internal/repository"
	"backend-rs100/internal/usecase"

	"cloud.google.com/go/firestore"
	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

// Simple helper to load .env file
func loadEnv(filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Printf("[WARN] Could not read %s file: %v", filename, err)
		return
	}
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			os.Setenv(key, value)
		}
	}
}

func main() {
	// 0. Load .env file into environment
	loadEnv("config/.env")

	// 1. Load Configurations
	cfg := config.LoadConfig()

	// 2. Initialize database connection (passing nil; in-memory fallback handles data access without a running MySQL instance)
	log.Printf("[INIT] Connecting to MySQL at %s:%s...", cfg.DBHost, cfg.DBPort)

	// Initialize Firestore
	ctx := context.Background()
	projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")
	if projectID == "" {
		log.Println("[WARN] GOOGLE_CLOUD_PROJECT is not set, trying firestore.DetectProjectID...")
		projectID = firestore.DetectProjectID
	}

	credsFile := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	var opts []option.ClientOption
	if credsFile != "" {
		log.Printf("[INIT] Using credentials file: %s", credsFile)
		opts = append(opts, option.WithCredentialsFile(credsFile))
	}

	firestoreClient, err := firestore.NewClient(ctx, projectID, opts...)
	if err != nil {
		log.Fatalf("[ERROR] Failed to create firestore client: %v", err)
	}
	defer firestoreClient.Close()
	log.Printf("[INIT] Connected to Firestore project: %s", projectID)

	// Initialize Storage Client
	storageClient, err := storage.NewClient(ctx, opts...)
	if err != nil {
		log.Fatalf("[ERROR] Failed to create storage client: %v", err)
	}
	defer storageClient.Close()
	
	bucketName := os.Getenv("GOOGLE_CLOUD_BUCKET")
	if bucketName == "" {
		bucketName = "rsas-storage-gallery"
		log.Printf("[WARN] GOOGLE_CLOUD_BUCKET is not set, using fallback: %s", bucketName)
	}
	log.Printf("[INIT] Using GCS bucket: %s", bucketName)

	// 3. Initialize Repositories
	bookRepo := repository.NewBookMySQLRepository(nil)
	activationRepo := repository.NewActivationMySQLRepository(nil)
	galleryRepo := repository.NewGalleryFirestoreRepository(firestoreClient)
	milestoneRepo := repository.NewMilestoneFirestoreRepository(firestoreClient)

	// 4. Initialize Usecases with 2 seconds context timeout
	timeout := 2 * time.Second
	bookUsecase := usecase.NewBookUsecase(bookRepo, timeout)
	authUsecase := usecase.NewAuthUsecase(activationRepo, timeout)
	galleryUsecase := usecase.NewGalleryUsecase(galleryRepo)
	milestoneUsecase := usecase.NewMilestoneUsecase(milestoneRepo)

	// 5. Initialize Delivery (HTTP Handler)
	handler := delivery.NewHttpHandler(bookUsecase, authUsecase, galleryUsecase, milestoneUsecase, storageClient, bucketName)

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
