package handlers

import (
	"click_counter/internal/service"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

type StatsHandler struct {
	statsService service.StatsService
}

func NewStatsHandler(statsService service.StatsService) *StatsHandler {
	return &StatsHandler{statsService: statsService}
}

func (h *StatsHandler) Handle(w http.ResponseWriter, r *http.Request) {
	bannerIDStr := r.URL.Path[len("/stats/"):]
	bannerID, err := strconv.Atoi(bannerIDStr)
	if err != nil {
		http.Error(w, "Invalid bannerID", http.StatusBadRequest)
		return
	}

	var req struct {
		From string `json:"tsFrom"`
		To   string `json:"tsTo"`
	}
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

	clicks, err := h.statsService.GetStatistics(bannerID, from, to)
	if err != nil {
		http.Error(w, "Failed to fetch statistics", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(strconv.Itoa(clicks) + "\n"))
}
