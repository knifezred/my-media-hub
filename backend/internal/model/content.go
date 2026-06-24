package model

import "time"

// ContentType 媒体内容类型
const (
	ContentTypePreview  = "preview"
	ContentTypeOCR      = "ocr"
	ContentTypeSummary  = "summary"
	ContentTypeSubtitle = "subtitle"
	ContentTypeNFO      = "nfo"
)

// MediaContent 媒体内容扩展
type MediaContent struct {
	MediaID     int64     `json:"media_id"`
	ContentType string    `json:"content_type"`
	Content     string    `json:"content"`
	UpdatedAt   time.Time `json:"updated_at"`
}
