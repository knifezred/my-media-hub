package model

import "time"

type Media struct {
	ID            int64      `json:"id"`
	MediaType     string     `json:"media_type"`
	Title         string     `json:"title"`
	Description   string     `json:"description"`
	Path          string     `json:"path"`
	Hash          string     `json:"hash"`
	Size          int64      `json:"size"`
	CoverPath     string     `json:"cover_path"`
	FavoriteCount int64      `json:"favorite_count"`
	ViewCount     int64      `json:"view_count"`
	RatingCount   int64      `json:"rating_count"`
	AvgRating     float64    `json:"avg_rating"`
	LastViewedAt  *time.Time `json:"last_viewed_at"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
}

type Tag struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type MediaTag struct {
	MediaID int64 `json:"media_id"`
	TagID   int64 `json:"tag_id"`
}

type Category struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	ParentID int64  `json:"parent_id"`
}

type MediaCategory struct {
	MediaID    int64 `json:"media_id"`
	CategoryID int64 `json:"category_id"`
}

type MediaMetadata struct {
	ID       int64  `json:"id"`
	MediaID  int64  `json:"media_id"`
	MetaKey  string `json:"meta_key"`
	MetaValue string `json:"meta_value"`
}
