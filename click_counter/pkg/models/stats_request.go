package models

type StatsRequest struct {
	From string `json:"tsFrom"`
	To   string `json:"tsTo"`
}
