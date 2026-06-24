package service

import (
	"database/sql"
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

func (s *BehaviorService) Record(mediaID int64, behaviorType string, score float64) error {
	_, err := s.behaviorRepo.Insert(mediaID, behaviorType, score)
	if err != nil {
		return err
	}

	switch behaviorType {
	case "favorite":
		count, _ := s.behaviorRepo.GetBehaviorCount(mediaID, "favorite")
		s.mediaRepo.UpdateStats(mediaID, count, 0, 0, 0)
	case "rating":
		count, _ := s.behaviorRepo.GetBehaviorCount(mediaID, "rating")
		avg, _ := s.behaviorRepo.GetAvgRating(mediaID)
		s.mediaRepo.UpdateStats(mediaID, 0, 0, count, avg)
	case "view":
		total, _ := s.behaviorRepo.GetBehaviorCount(mediaID, "view")
		s.mediaRepo.UpdateViewStats(mediaID, total)
	}

	return nil
}

func (s *BehaviorService) RemoveBehavior(mediaID int64, behaviorType string) error {
	return s.behaviorRepo.DeleteByMediaIDAndType(mediaID, behaviorType)
}

func (s *BehaviorService) IsFavorited(mediaID int64) (bool, error) {
	count, err := s.behaviorRepo.GetBehaviorCount(mediaID, "favorite")
	return count > 0, err
}

func (s *BehaviorService) GetRating(mediaID int64) (int, error) {
	score, err := s.behaviorRepo.GetMaxScoreByType(mediaID, "rating")
	return int(score), err
}

func (s *BehaviorService) IsViewed(mediaID int64) (bool, error) {
	count, err := s.behaviorRepo.GetBehaviorCount(mediaID, "view")
	return count > 0, err
}

func (s *BehaviorService) IsHidden(mediaID int64) (bool, error) {
	count, err := s.behaviorRepo.GetBehaviorCount(mediaID, "hidden")
	return count > 0, err
}

func (s *BehaviorService) ListFavorites(req model.FavoritePageRequest) (*model.PageResponse, error) {
	items, total, err := s.behaviorRepo.ListByType("favorite", req.Page, req.PageSize)
	if err != nil {
		return nil, err
	}
	page := req.Page
	if page < 1 {
		page = 1
	}
	pageSize := req.PageSize
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}
	return &model.PageResponse{Items: items, Total: total, Page: page, PageSize: pageSize}, nil
}

func (s *BehaviorService) ListViewed(page, pageSize int) (*model.PageResponse, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}
	items, total, err := s.behaviorRepo.ListByType("view", page, pageSize)
	if err != nil {
		return nil, err
	}
	return &model.PageResponse{Items: items, Total: total, Page: page, PageSize: pageSize}, nil
}

func (s *BehaviorService) GetStatistics() (*model.BehaviorStatistics, error) {
	return s.behaviorRepo.GetBehaviorStatistics()
}

func (s *BehaviorService) ListFavoritesAsMedia(page, pageSize int) (*model.PageResponse, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}
	items, total, err := s.behaviorRepo.ListMediaByType("favorite", page, pageSize)
	if err != nil {
		return nil, err
	}
	return &model.PageResponse{Items: items, Total: total, Page: page, PageSize: pageSize}, nil
}

func (s *BehaviorService) ListViewedAsMedia(page, pageSize int) (*model.PageResponse, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}
	items, total, err := s.behaviorRepo.ListMediaByType("view", page, pageSize)
	if err != nil {
		return nil, err
	}
	return &model.PageResponse{Items: items, Total: total, Page: page, PageSize: pageSize}, nil
}

func (s *BehaviorService) ListHiddenAsMedia(page, pageSize int) (*model.PageResponse, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}
	items, total, err := s.behaviorRepo.ListMediaByType("hidden", page, pageSize)
	if err != nil {
		return nil, err
	}
	return &model.PageResponse{Items: items, Total: total, Page: page, PageSize: pageSize}, nil
}
