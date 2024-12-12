package handlers

import (
	"click_counter/internal/service"
	"click_counter/pkg/models"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

type StatsHandler struct {
	service *service.ClickService
}

func NewStatsHandler(service *service.ClickService) *StatsHandler {
	return &StatsHandler{service: service}
}

func (h *StatsHandler) Handle(w http.ResponseWriter, r *http.Request) {
	bannerIDStr := r.URL.Path[len("/stats/"):]
	bannerID, err := strconv.Atoi(bannerIDStr)
	if err != nil {
		http.Error(w, "Invalid bannerID", http.StatusBadRequest)
		return
	}

	var req models.StatsRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	from, err := time.Parse(time.RFC3339, req.From)
	if err != nil {
		http.Error(w, "Invalid tsFrom format", http.StatusBadRequest)
		return
	}

	to, err := time.Parse(time.RFC3339, req.To)
	if err != nil {
		http.Error(w, "Invalid tsTo format", http.StatusBadRequest)
		return
	}

	total := h.service.GetStatistics(bannerID, from, to)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(strconv.Itoa(total) + "\n"))
}
