package main

import (
	"log"
	"net/http"
	"video-ads-backend/config"
	"video-ads-backend/internal/ads"
	"video-ads-backend/internal/storage"
	"video-ads-backend/internal/utils"
	"video-ads-backend/routes"
)

func main() {
	// Load env vars
	config.LoadEnv()

	// Logger and DB init
	logger := utils.NewLogger()
	db := storage.NewPostgresDB()

	// Init ads module
	adRepo := ads.NewRepository(db)
	adService := ads.NewService(adRepo)
	adHandler := ads.NewHandler(adService, logger)

	// Register routes
	r := routes.RegisterRoutes(adHandler, logger)

	log.Println("🚀 Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
