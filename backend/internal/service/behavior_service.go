package service

import (
	"database/sql"
	"fmt"
	"my-media-hub/backend/internal/model"
	"my-media-hub/backend/internal/repository"
)

type BehaviorService struct {
	behaviorRepo *repository.BehaviorRepository
	mediaRepo    *repository.MediaRepository
}

func NewBehaviorService(db *sql.DB) *BehaviorService {
	return &BehaviorService{
		behaviorRepo: repository.NewBehaviorRepository(db),
		mediaRepo:    repository.NewMediaRepository(db),
	}
}

// Record 记录行为 + 同步 media 状态
func (s *BehaviorService) Record(mediaID int64, behaviorType, behaviorValue, behaviorSource string) error {
	if _, err := s.behaviorRepo.Insert(mediaID, behaviorType, behaviorValue, behaviorSource); err != nil {
		return err
	}

	// 同步 media 当前状态字段（双写：状态+历史）
	switch behaviorType {
	case model.BehaviorFavorite:
		return s.mediaRepo.SetFavorite(mediaID, true)
	case model.BehaviorUnfavorite:
		return s.mediaRepo.SetFavorite(mediaID, false)
	case model.BehaviorRate:
		// behaviorValue = {"rating": 4.5}
		var rating float64
		if err := parseRatingValue(behaviorValue, &rating); err != nil {
			return err
		}
		return s.mediaRepo.SetRating(mediaID, rating)
	case model.BehaviorHide:
		return s.mediaRepo.SetHidden(mediaID, true)
	case model.BehaviorUnhide:
		return s.mediaRepo.SetHidden(mediaID, false)
	case model.BehaviorView:
		return s.mediaRepo.IncViewCount(mediaID)
	}
	return nil
}

func (s *BehaviorService) Favorite(mediaID int64) error {
	return s.Record(mediaID, model.BehaviorFavorite, "{}", model.BehaviorSourceManual)
}

func (s *BehaviorService) Unfavorite(mediaID int64) error {
	return s.Record(mediaID, model.BehaviorUnfavorite, "{}", model.BehaviorSourceManual)
}

func (s *BehaviorService) Rate(mediaID int64, rating float64) error {
	if rating < 0.5 || rating > 5.0 {
		return fmt.Errorf("rating out of range [0.5, 5.0]: %v", rating)
	}
	value := fmt.Sprintf(`{"rating":%.1f}`, rating)
	return s.Record(mediaID, model.BehaviorRate, value, model.BehaviorSourceManual)
}

func (s *BehaviorService) Hide(mediaID int64) error {
	return s.Record(mediaID, model.BehaviorHide, "{}", model.BehaviorSourceManual)
}

func (s *BehaviorService) Unhide(mediaID int64) error {
	return s.Record(mediaID, model.BehaviorUnhide, "{}", model.BehaviorSourceManual)
}

func (s *BehaviorService) View(mediaID int64) error {
	return s.Record(mediaID, model.BehaviorView, "{}", model.BehaviorSourceManual)
}

func (s *BehaviorService) IsFavorited(mediaID int64) (bool, error) {
	m, err := s.mediaRepo.GetByID(mediaID)
	if err != nil || m == nil {
		return false, err
	}
	return m.Favorite, nil
}

func (s *BehaviorService) IsHidden(mediaID int64) (bool, error) {
	m, err := s.mediaRepo.GetByID(mediaID)
	if err != nil || m == nil {
		return false, err
	}
	return m.Hidden, nil
}

func (s *BehaviorService) GetStatistics() (*model.BehaviorStatistics, error) {
	return s.behaviorRepo.Statistics()
}

// parseRatingValue 从 {"rating":4.5} 中提取评分值
func parseRatingValue(jsonStr string, rating *float64) error {
	var r float64
	if _, err := fmt.Sscanf(jsonStr, `{"rating":%f}`, &r); err != nil {
		return fmt.Errorf("parse rating value: %w", err)
	}
	*rating = r
	return nil
}
