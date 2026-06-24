package model

import "time"

type SearchHistory struct {
	ID         int64     `json:"id"`
	Keyword    string    `json:"keyword"`
	KeywordNorm string   `json:"keyword_norm"`
	UseCount   int       `json:"use_count"`
	LastUsedAt time.Time `json:"last_used_at"`
	CreatedAt  time.Time `json:"created_at"`
}

type ScannerIndex struct {
	ID           int64     `json:"id"`
	MediaID      *int64    `json:"media_id"`
	FilePath     string    `json:"file_path"`
	FileSize     int64     `json:"file_size"`
	ModifiedTime time.Time `json:"modified_time"`
	FileHash     string    `json:"file_hash"`
	LastScanAt   time.Time `json:"last_scan_at"`
}
