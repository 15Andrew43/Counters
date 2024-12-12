package repository

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type Repository interface {
	AddClick(bannerID int) error
	GetClicks(bannerID int, from, to time.Time) (int, error)
}

type ClickRepository struct {
	db *sql.DB
}

func NewClickRepository(host, port, user, password, dbname string) (*ClickRepository, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Printf("Error opening database connection: %v", err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		log.Printf("Error pinging database: %v", err)
		return nil, err
	}

	log.Println("Connected to the database successfully")
	return &ClickRepository{db: db}, nil
}

func (r *ClickRepository) AddClick(bannerID int) error {
	query := `
		INSERT INTO clicks (banner_id, timestamp, count)
		VALUES ($1, $2, $3)
	`
	_, err := r.db.Exec(query, bannerID, time.Now(), 1)
	if err != nil {
		log.Printf("Error inserting click for bannerID %d: %v", bannerID, err)
		return err
	}

	log.Printf("Click added for bannerID %d", bannerID)
	return nil
}

func (r *ClickRepository) GetClicks(bannerID int, from, to time.Time) (int, error) {
	query := `
		SELECT COALESCE(SUM(count), 0)
		FROM clicks
		WHERE banner_id = $1 AND timestamp BETWEEN $2 AND $3
	`
	var total int
	err := r.db.QueryRow(query, bannerID, from, to).Scan(&total)
	if err != nil {
		log.Printf("Error fetching clicks for bannerID %d: %v", bannerID, err)
		return 0, err
	}

	log.Printf("Fetched %d clicks for bannerID %d from %s to %s", total, bannerID, from, to)
	return total, nil
}
