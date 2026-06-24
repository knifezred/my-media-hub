package model

type PageRequest struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

type PageResponse struct {
	Items    interface{} `json:"items"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"page_size"`
}

type MediaPageRequest struct {
	Page       int     `json:"page"`
	PageSize   int     `json:"page_size"`
	MediaType  string  `json:"media_type"`
	CategoryID int64   `json:"category_id"`
	TagID      int64   `json:"tag_id"`
	Sort       string  `json:"sort"`
	TagIDs     []int64 `json:"tag_ids"`
}

type MediaDetail struct {
	Media
	Tags       []Tag     `json:"tags"`
	Categories []Category `json:"categories"`
	Metadata   map[string]string `json:"metadata"`
	Favorite   bool      `json:"favorite"`
	Rating     int       `json:"rating"`
	Viewed     bool      `json:"viewed"`
	Hidden     bool      `json:"hidden"`
}

type SearchRequest struct {
	Keyword   string `json:"keyword"`
	MediaType string `json:"media_type"`
	Page      int    `json:"page"`
	PageSize  int    `json:"page_size"`
}

type FavoritePageRequest struct {
	Page      int    `json:"page"`
	PageSize  int    `json:"page_size"`
	MediaType string `json:"media_type"`
}

type BehaviorRequest struct {
	MediaID      int64   `json:"media_id"`
	BehaviorType string  `json:"behavior_type"`
	Score        float64 `json:"score"`
}

type BehaviorStatistics struct {
	FavoriteCount int64 `json:"favorite_count"`
	ViewCount     int64 `json:"view_count"`
	RatingCount   int64 `json:"rating_count"`
	HiddenCount   int64 `json:"hidden_count"`
}

type StatsOverview struct {
	TotalMedia    int64 `json:"total_media"`
	TotalImages   int64 `json:"total_images"`
	TotalVideos   int64 `json:"total_videos"`
	TotalNovels   int64 `json:"total_novels"`
	FavoriteCount int64 `json:"favorite_count"`
	ViewedCount   int64 `json:"viewed_count"`
}

type ScannerStatus struct {
	Running   bool    `json:"running"`
	Processed int     `json:"processed"`
	Total     int     `json:"total"`
	Progress  float64 `json:"progress"`
}

type RateRequest struct {
	MediaID int64 `json:"media_id"`
	Rating  int   `json:"rating"`
}

type MediaIDRequest struct {
	MediaID int64 `json:"media_id"`
}
