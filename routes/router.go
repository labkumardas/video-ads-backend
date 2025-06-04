package routes

import (
	"net/http"
	"video-ads-backend/internal/ads"

	"go.uber.org/zap"

	"github.com/gorilla/mux"
)

func RegisterRoutes(adHandler *ads.Handler, logger *zap.Logger) http.Handler {
	r := mux.NewRouter()

	// API Endpoints
	r.HandleFunc("/ads", adHandler.GetAds).Methods("GET")
	r.HandleFunc("/ads/click", adHandler.PostClick).Methods("POST")
	r.HandleFunc("/ads/analytics", adHandler.GetAnalytics).Methods("GET")

	// Healthcheck endpoint
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	return r
}
