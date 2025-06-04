package ads

import (
	"encoding/json"
	"net/http"
	"video-ads-backend/internal/utils"

	"go.uber.org/zap"
)

type Handler struct {
	service *Service
	logger  *zap.Logger
}

func NewHandler(s *Service, logger *zap.Logger) *Handler {
	return &Handler{s, logger}
}

func (h *Handler) GetAds(w http.ResponseWriter, r *http.Request) {
	ads, err := h.service.FetchAds()
	if err != nil {
		h.logger.Error("fetch ads error", zap.Error(err))
		utils.ErrorResponse(w, "failed to fetch ads", 500)
		return
	}
	utils.SuccessResponse(w, ads)
}

func (h *Handler) PostClick(w http.ResponseWriter, r *http.Request) {
	var click Click
	if err := json.NewDecoder(r.Body).Decode(&click); err != nil {
		utils.ErrorResponse(w, "invalid input", 400)
		return
	}
	go h.service.LogClick(click) // Non-blocking!
	utils.SuccessResponse(w, map[string]string{"status": "click received"})
}

func (h *Handler) GetAnalytics(w http.ResponseWriter, r *http.Request) {
	data, err := h.service.GetClickStats()
	if err != nil {
		utils.ErrorResponse(w, "failed to get analytics", 500)
		return
	}
	utils.SuccessResponse(w, data)
}
