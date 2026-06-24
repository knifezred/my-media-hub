package model

import "time"

type SearchHistory struct {
	ID           int64     `json:"id"`
	Keyword      string    `json:"keyword"`
	SearchSource string    `json:"search_source"`
	ResultCount  int       `json:"result_count"`
	CreatedAt    time.Time `json:"created_at"`
}
