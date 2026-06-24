package model

import "time"

// BehaviorType 用户行为类型
const (
	BehaviorView       = "view"
	BehaviorFavorite   = "favorite"
	BehaviorUnfavorite = "unfavorite"
	BehaviorRate       = "rate"
	BehaviorHide       = "hide"
	BehaviorUnhide     = "unhide"
)

// BehaviorSource 行为来源
const (
	BehaviorSourceManual         = "manual"
	BehaviorSourceSearch         = "search"
	BehaviorSourceRecommendation = "recommendation"
	BehaviorSourceHomeFeed       = "home_feed"
)

// MediaBehavior 用户行为流水
type MediaBehavior struct {
	ID             int64     `json:"id"`
	MediaID        int64     `json:"media_id"`
	BehaviorType   string    `json:"behavior_type"`
	BehaviorValue  string    `json:"behavior_value"`
	BehaviorSource string    `json:"behavior_source"`
	CreatedAt      time.Time `json:"created_at"`
}
