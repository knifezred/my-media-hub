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
	Page      int    `json:"page"`
	PageSize  int    `json:"page_size"`
	MediaType string `json:"media_type"`
	Status    string `json:"status"`
	Sort      string `json:"sort"`
}

type MediaDetail struct {
	Media
	Tags       []Tag      `json:"tags"`
	Categories []Category `json:"categories"`
}

type SearchRequest struct {
	Keyword   string `json:"keyword"`
	MediaType string `json:"media_type"`
	Page      int    `json:"page"`
	PageSize  int    `json:"page_size"`
}

// BehaviorRequest 记录行为请求（v2.1）
type BehaviorRequest struct {
	MediaID        int64  `json:"media_id"`
	BehaviorType   string `json:"behavior_type"`
	BehaviorValue  string `json:"behavior_value"`
	BehaviorSource string `json:"behavior_source"`
}

// RateRequest 评分请求（v2.1 支持 0.5 步进）
type RateRequest struct {
	MediaID int64   `json:"media_id"`
	Rating  float64 `json:"rating"`
}

type MediaIDRequest struct {
	MediaID int64 `json:"media_id"`
}

type BehaviorStatistics struct {
	FavoriteCount int64 `json:"favorite_count"`
	ViewCount     int64 `json:"view_count"`
	RateCount     int64 `json:"rate_count"`
	HideCount     int64 `json:"hide_count"`
}

type StatsOverview struct {
	TotalMedia  int64 `json:"total_media"`
	TotalImages int64 `json:"total_images"`
	TotalVideos int64 `json:"total_videos"`
	TotalNovels int64 `json:"total_novels"`
	TotalMusic  int64 `json:"total_music"`
}

type ScannerStatus struct {
	Running   bool    `json:"running"`
	Processed int     `json:"processed"`
	Total     int     `json:"total"`
	Progress  float64 `json:"progress"`
}
