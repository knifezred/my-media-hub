package service

import (
	"database/sql"
	"my-media-hub/backend/internal/model"
	"my-media-hub/backend/internal/repository"
	"my-media-hub/backend/internal/search"
)

type SearchService struct {
	search  *repository.SearchRepository
	history *repository.SearchHistoryRepository
}

func NewSearchService(db *sql.DB, index *search.Index) *SearchService {
	return &SearchService{
		search:  repository.NewSearchRepository(db, index),
		history: repository.NewSearchHistoryRepository(db),
	}
}

func (s *SearchService) Search(req model.SearchRequest) (*model.PageResponse, error) {
	items, total, err := s.search.Search(req.Keyword, req.MediaType, req.Page, req.PageSize)
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
	// 记录搜索历史
	s.history.InsertOrUpdate(req.Keyword)
	return &model.PageResponse{
		Items:    items,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	}, nil
}

func (s *SearchService) Suggestions(prefix string) ([]string, error) {
	return s.search.Suggestions(prefix, 10)
}

func (s *SearchService) ListHistory(page, pageSize int) (*model.PageResponse, error) {
	items, total, err := s.history.List(page, pageSize)
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

func (s *SearchService) DeleteHistory(id int64) error {
	return s.history.Delete(id)
}

func (s *SearchService) ClearHistory() error {
	return s.history.Clear()
}
