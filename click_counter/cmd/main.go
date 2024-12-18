package main

import (
	"click_counter/internal/handlers"
	"click_counter/internal/repository"
	"click_counter/internal/service"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	flushIntervalStr := os.Getenv("FLUSH_INTERVAL")

	flushInterval, err := strconv.Atoi(flushIntervalStr)
	if err != nil || flushInterval <= 0 {
		log.Fatalf("Invalid flush interval: %v", err)
	}

	repo, err := repository.NewClickRepository(host, port, user, password, dbname)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	clickService := service.NewClickService(repo, time.Duration(flushInterval)*time.Second)
	statsService := service.NewStatsService(repo)

	counterHandler := handlers.NewCounterHandler(clickService)
	statsHandler := handlers.NewStatsHandler(statsService)

	http.HandleFunc("/counter/", counterHandler.Handle)
	http.HandleFunc("/stats/", statsHandler.Handle)

	port = ":8080"
	log.Printf("Server is running on port %s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
