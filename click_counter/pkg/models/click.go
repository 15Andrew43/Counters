package models

import "time"

type Click struct {
	// BannerID  int
	Timestamp time.Time
	Count     int
}
