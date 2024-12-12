package service

import (
	"click_counter/internal/repository"
	"log"
	"time"
)

type StatsService interface {
	GetStatistics(bannerID int, from, to time.Time) (int, error)
}

type DefaultStatsService struct {
	repo repository.Repository
}

func NewStatsService(repo repository.Repository) *DefaultStatsService {
	return &DefaultStatsService{repo: repo}
}

func (s *DefaultStatsService) GetStatistics(bannerID int, from, to time.Time) (int, error) {
	clicks, err := s.repo.GetClicks(bannerID, from, to)
	if err != nil {
		log.Printf("Error fetching statistics for bannerID %d: %v", bannerID, err)
		return 0, err
	}

	log.Printf("Fetched %d clicks for bannerID %d from %s to %s", clicks, bannerID, from, to)
	return clicks, nil
}
