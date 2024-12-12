package service

import (
	"click_counter/internal/repository"
	"time"
)

type Service interface {
	RegisterClick(bannerID int)
	GetStatistics(bannerID int, from, to time.Time) int
}

type ClickService struct {
	repo repository.Repository
}

func NewClickService(repo repository.Repository) *ClickService {
	return &ClickService{repo: repo}
}

func (s *ClickService) RegisterClick(bannerID int) {
	s.repo.AddClick(bannerID)
}

func (s *ClickService) GetStatistics(bannerID int, from, to time.Time) int {
	return s.repo.GetClicks(bannerID, from, to)
}
