package service

import (
	"database/sql"
	"my-media-hub/backend/internal/model"
	"my-media-hub/backend/internal/repository"
)

type MediaService struct {
	media        *repository.MediaRepository
	tagRepo      *repository.TagRepository
	categoryRepo *repository.CategoryRepository
	mediaTag     *repository.MediaTagRepository
	mediaCat     *repository.MediaCategoryRepository
	mediaMeta    *repository.MediaMetadataRepository
	behaviorSvc  *BehaviorService
}

func NewMediaService(db *sql.DB, behaviorSvc *BehaviorService) *MediaService {
	return &MediaService{
		media:        repository.NewMediaRepository(db),
		tagRepo:      repository.NewTagRepository(db),
		categoryRepo: repository.NewCategoryRepository(db),
		mediaTag:     repository.NewMediaTagRepository(db),
		mediaCat:     repository.NewMediaCategoryRepository(db),
		mediaMeta:    repository.NewMediaMetadataRepository(db),
		behaviorSvc:  behaviorSvc,
	}
}

func (s *MediaService) GetByID(id int64) (*model.MediaDetail, error) {
	m, err := s.media.GetByID(id)
	if err != nil {
		return nil, err
	}
	if m == nil {
		return nil, nil
	}

	tags, _ := s.tagRepo.GetByMediaID(id)
	categories, _ := s.categoryRepo.GetByMediaID(id)
	metadata, _ := s.mediaMeta.GetByMediaID(id)

	favorite, _ := s.behaviorSvc.IsFavorited(id)
	rating, _ := s.behaviorSvc.GetRating(id)
	viewed, _ := s.behaviorSvc.IsViewed(id)
	hidden, _ := s.behaviorSvc.IsHidden(id)

	detail := &model.MediaDetail{
		Media:      *m,
		Tags:       tags,
		Categories: categories,
		Metadata:   metadata,
		Favorite:   favorite,
		Rating:     rating,
		Viewed:     viewed,
		Hidden:     hidden,
	}

	return detail, nil
}

func (s *MediaService) List(req model.MediaPageRequest) (*model.PageResponse, error) {
	items, total, err := s.media.List(req)
	if err != nil {
		return nil, err
	}

	if req.Page < 1 {
		req.Page = 1
	}
	pageSize := req.PageSize
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}

	return &model.PageResponse{
		Items:    items,
		Total:    total,
		Page:     req.Page,
		PageSize: pageSize,
	}, nil
}

func (s *MediaService) GetTagsByMediaID(mediaID int64) ([]model.Tag, error) {
	return s.tagRepo.GetByMediaID(mediaID)
}

func (s *MediaService) GetCategoriesByMediaID(mediaID int64) ([]model.Category, error) {
	return s.categoryRepo.GetByMediaID(mediaID)
}

func (s *MediaService) ListTags(page, pageSize int) (*model.PageResponse, error) {
	items, total, err := s.tagRepo.List(page, pageSize)
	if err != nil {
		return nil, err
	}
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}
	return &model.PageResponse{Items: items, Total: total, Page: page, PageSize: pageSize}, nil
}

func (s *MediaService) ListCategories(page, pageSize int) (*model.PageResponse, error) {
	items, total, err := s.categoryRepo.List(page, pageSize)
	if err != nil {
		return nil, err
	}
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}
	return &model.PageResponse{Items: items, Total: total, Page: page, PageSize: pageSize}, nil
}

func (s *MediaService) GetTagByID(id int64) (*model.Tag, error) {
	return s.tagRepo.GetByID(id)
}
