package main

import (
	"click_counter/internal/handlers"
	"click_counter/internal/repository"
	"click_counter/internal/service"
	"log"
	"net/http"
)

func main() {

	repo := repository.NewClickRepository()

	svc := service.NewClickService(repo)

	counterHandler := handlers.NewCounterHandler(svc)
	statsHandler := handlers.NewStatsHandler(svc)

	http.HandleFunc("/counter/", counterHandler.Handle)
	http.HandleFunc("/stats/", statsHandler.Handle)

	port := ":8080"
	log.Printf("Server is running on port %s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
