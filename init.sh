#!/bin/bash

# Создаем структуру директорий
mkdir -p click_counter/{cmd,internal/{handlers,repository,service},pkg/models}

# Переходим в папку проекта
cd click_counter || exit

# Инициализация Go-модуля
go mod init click_counter

# Заполнение файлов

# Создание файла модели
cat <<EOL > pkg/models/click.go
package models

type Click struct {
    BannerID  int
    Timestamp int64
    Count     int
}
EOL

# Создание файла репозитория
cat <<EOL > internal/repository/repository.go
package repository

import (
    "sync"
    "click_counter/pkg/models"
)

type ClickRepository struct {
    mu     sync.Mutex
    clicks map[int]int
}

func NewClickRepository() *ClickRepository {
    return &ClickRepository{
        clicks: make(map[int]int),
    }
}

func (r *ClickRepository) AddClick(bannerID int) {
    r.mu.Lock()
    defer r.mu.Unlock()
    r.clicks[bannerID]++
}

func (r *ClickRepository) GetClicks(bannerID int) int {
    r.mu.Lock()
    defer r.mu.Unlock()
    return r.clicks[bannerID]
}
EOL

# Создание файла сервиса
cat <<EOL > internal/service/service.go
package service

import (
    "click_counter/internal/repository"
)

type ClickService struct {
    repo *repository.ClickRepository
}

func NewClickService(repo *repository.ClickRepository) *ClickService {
    return &ClickService{repo: repo}
}

func (s *ClickService) RegisterClick(bannerID int) {
    s.repo.AddClick(bannerID)
}

func (s *ClickService) GetStatistics(bannerID int) int {
    return s.repo.GetClicks(bannerID)
}
EOL

# Создание файла хендлеров
cat <<EOL > internal/handlers/handlers.go
package handlers

import (
    "net/http"
    "strconv"
    "click_counter/internal/service"
)

type ClickHandler struct {
    service *service.ClickService
}

func NewClickHandler(service *service.ClickService) *ClickHandler {
    return &ClickHandler{service: service}
}

func (h *ClickHandler) CounterHandler(w http.ResponseWriter, r *http.Request) {
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

func (h *ClickHandler) StatsHandler(w http.ResponseWriter, r *http.Request) {
    bannerIDStr := r.URL.Path[len("/stats/"):]
    bannerID, err := strconv.Atoi(bannerIDStr)
    if err != nil {
        http.Error(w, "Invalid bannerID", http.StatusBadRequest)
        return
    }

    clicks := h.service.GetStatistics(bannerID)
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(strconv.Itoa(clicks) + "\n"))
}
EOL

# Создание файла main.go
cat <<EOL > cmd/main.go
package main

import (
    "log"
    "net/http"
    "click_counter/internal/handlers"
    "click_counter/internal/repository"
    "click_counter/internal/service"
)

func main() {
    repo := repository.NewClickRepository()
    service := service.NewClickService(repo)
    handler := handlers.NewClickHandler(service)

    http.HandleFunc("/counter/", handler.CounterHandler)
    http.HandleFunc("/stats/", handler.StatsHandler)

    port := ":8080"
    log.Printf("Server is running on port %s\n", port)
    if err := http.ListenAndServe(port, nil); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}
EOL

echo "Структура проекта создана."
tree .
