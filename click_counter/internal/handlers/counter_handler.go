package handlers

import (
	"click_counter/internal/service"
	"net/http"
	"strconv"
)

type CounterHandler struct {
	service *service.ClickService
}

func NewCounterHandler(service *service.ClickService) *CounterHandler {
	return &CounterHandler{service: service}
}

func (h *CounterHandler) Handle(w http.ResponseWriter, r *http.Request) {
	bannerIDStr := r.URL.Path[len("/counter/"):]
	bannerID, err := strconv.Atoi(bannerIDStr)
	if err != nil {
		http.Error(w, "Invalid bannerID", http.StatusBadRequest)
		return
	}

	h.service.RegisterClick(bannerID)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Click registered\n"))
}
