package repository

import (
	"click_counter/pkg/models"
	"sync"
	"time"
)

type Repository interface {
	AddClick(bannerID int)
	GetClicks(bannerID int, from, to time.Time) int
}

type ClickRepository struct {
	clicks map[int][]models.Click
	mu     sync.Mutex
}

func NewClickRepository() *ClickRepository {
	return &ClickRepository{
		clicks: make(map[int][]models.Click),
	}
}

func (r *ClickRepository) AddClick(bannerID int) {
	r.mu.Lock()
	defer r.mu.Unlock()

	now := time.Now()
	r.clicks[bannerID] = append(r.clicks[bannerID], models.Click{Timestamp: now, Count: 1})
}

func (r *ClickRepository) GetClicks(bannerID int, from, to time.Time) int {
	r.mu.Lock()
	defer r.mu.Unlock()

	total := 0
	if clicks, exists := r.clicks[bannerID]; exists {
		for _, click := range clicks {
			if click.Timestamp.After(from) && click.Timestamp.Before(to) {
				total += click.Count
			}
		}
	}
	return total
}
