package service

import (
	"click_counter/internal/repository"
	"log"
	"sync"
	"time"
)

type ClickService interface {
	RegisterClick(bannerID int) error
}

type DefaultClickService struct {
	repo          repository.Repository
	clicksCache   map[int]int
	cacheMutex    sync.Mutex
	flushInterval time.Duration
	stopFlush     chan struct{}
}

func NewClickService(repo repository.Repository, flushInterval time.Duration) *DefaultClickService {
	service := &DefaultClickService{
		repo:          repo,
		clicksCache:   make(map[int]int),
		flushInterval: flushInterval,
		stopFlush:     make(chan struct{}),
	}

	go service.periodicFlush()
	return service
}

func (s *DefaultClickService) RegisterClick(bannerID int) error {
	s.cacheMutex.Lock()
	defer s.cacheMutex.Unlock()
	s.clicksCache[bannerID]++
	return nil
}

func (s *DefaultClickService) periodicFlush() {
	ticker := time.NewTicker(s.flushInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			s.flushToDB()
		case <-s.stopFlush:
			return
		}
	}
}

func (s *DefaultClickService) flushToDB() {
	s.cacheMutex.Lock()
	defer s.cacheMutex.Unlock()

	now := time.Now()
	roundedTimestamp := time.Unix(now.Unix()-now.Unix()%int64(s.flushInterval.Seconds()), 0)

	for bannerID, count := range s.clicksCache {
		if err := s.repo.AddClicks(bannerID, count, roundedTimestamp); err != nil {
			log.Printf("Error flushing clicks for bannerID %d: %v", bannerID, err)
			continue
		}
	}
}

func (s *DefaultClickService) Stop() {
	close(s.stopFlush)
	s.flushToDB()
}
