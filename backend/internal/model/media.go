package model

import "time"

// MediaStatus 媒体生命周期状态
const (
	MediaStatusNew     = "new"
	MediaStatusParsing = "parsing"
	MediaStatusActive  = "active"
	MediaStatusMissing = "missing"
	MediaStatusError   = "error"
	MediaStatusDeleted = "deleted"
)

// MediaType 媒体类型
const (
	MediaTypeImage = "image"
	MediaTypeVideo = "video"
	MediaTypeNovel = "novel"
	MediaTypeMusic = "music"
)

// Media 媒体主表
type Media struct {
	ID              int64      `json:"id"`
	MediaType       string     `json:"media_type"`
	Title           string     `json:"title"`
	Description     string     `json:"description"`
	Path            string     `json:"path"`
	Hash            string     `json:"hash"`
	Size            int64      `json:"size"`
	CoverPath       string     `json:"cover_path"`
	Status          string     `json:"status"`
	LastError       string     `json:"last_error"`
	MetadataJSON    string     `json:"metadata_json"`
	MetadataVersion int        `json:"metadata_version"`
	Favorite        bool       `json:"favorite"`
	FavoriteAt      *time.Time `json:"favorite_at"`
	Rating          float64    `json:"rating"`
	RatingAt        *time.Time `json:"rating_at"`
	Hidden          bool       `json:"hidden"`
	HiddenAt        *time.Time `json:"hidden_at"`
	ViewCount       int        `json:"view_count"`
	LastViewedAt    *time.Time `json:"last_viewed_at"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
}

// Tag 标签字典
type Tag struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	NameNorm  string    `json:"name_norm"`
	Source    string    `json:"source"`
	CreatedAt time.Time `json:"created_at"`
}

// Category 分类字典
type Category struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	ParentID  *int64    `json:"parent_id"`
	Level     int       `json:"level"`
	Path      string    `json:"path"`
	Sort      int       `json:"sort"`
	CreatedAt time.Time  `json:"created_at"`
}
