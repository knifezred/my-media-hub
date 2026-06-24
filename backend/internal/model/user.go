package model

import "time"

type UserBehavior struct {
	ID           int64     `json:"id"`
	MediaID      int64     `json:"media_id"`
	BehaviorType string    `json:"behavior_type"`
	Score        float64   `json:"score"`
	CreatedAt    time.Time `json:"created_at"`
}
