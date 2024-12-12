package service

import (
	"click_counter/internal/repository"
	"log"
)

type ClickService interface {
	RegisterClick(bannerID int) error
}

type DefaultClickService struct {
	repo repository.Repository
}

func NewClickService(repo repository.Repository) *DefaultClickService {
	return &DefaultClickService{repo: repo}
}

func (s *DefaultClickService) RegisterClick(bannerID int) error {
	err := s.repo.AddClick(bannerID)
	if err != nil {
		log.Printf("Error registering click for bannerID %d: %v", bannerID, err)
		return err
	}

	log.Printf("Click registered for bannerID %d", bannerID)
	return nil
}
