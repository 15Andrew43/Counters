package handlers

import (
	"click_counter/internal/service"
	"net/http"
	"strconv"
)

type CounterHandler struct {
	clickService service.ClickService
}

func NewCounterHandler(clickService service.ClickService) *CounterHandler {
	return &CounterHandler{clickService: clickService}
}

func (h *CounterHandler) Handle(w http.ResponseWriter, r *http.Request) {
	bannerIDStr := r.URL.Path[len("/counter/"):]
	bannerID, err := strconv.Atoi(bannerIDStr)
	if err != nil {
		http.Error(w, "Invalid bannerID", http.StatusBadRequest)
		return
	}

	err = h.clickService.RegisterClick(bannerID)
	if err != nil {
		http.Error(w, "Failed to register click", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Click registered\n"))
}
