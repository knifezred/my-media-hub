package repository

import (
	"database/sql"
	"fmt"
	"my-media-hub/backend/internal/model"
)

type SearchHistoryRepository struct {
	db *sql.DB
}

func NewSearchHistoryRepository(db *sql.DB) *SearchHistoryRepository {
	return &SearchHistoryRepository{db: db}
}

func (r *SearchHistoryRepository) Insert(keyword, searchSource string, resultCount int) (int64, error) {
	result, err := r.db.Exec(
		"INSERT INTO search_history (keyword, search_source, result_count) VALUES (?, ?, ?)",
		keyword, searchSource, resultCount,
	)
	if err != nil {
		return 0, fmt.Errorf("insert search history: %w", err)
	}
	return result.LastInsertId()
}

func (r *SearchHistoryRepository) List(page, pageSize int) ([]model.SearchHistory, int64, error) {
	var total int64
	err := r.db.QueryRow("SELECT COUNT(*) FROM search_history").Scan(&total)
	if err != nil {
		return nil, 0, fmt.Errorf("count search history: %w", err)
	}

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}
	offset := (page - 1) * pageSize

	rows, err := r.db.Query(
		"SELECT id, keyword, search_source, result_count, created_at FROM search_history ORDER BY created_at DESC LIMIT ? OFFSET ?",
		pageSize, offset,
	)
	if err != nil {
		return nil, 0, fmt.Errorf("list search history: %w", err)
	}
	defer rows.Close()

	items := make([]model.SearchHistory, 0)
	for rows.Next() {
		var h model.SearchHistory
		if err := rows.Scan(&h.ID, &h.Keyword, &h.SearchSource, &h.ResultCount, &h.CreatedAt); err != nil {
			return nil, 0, fmt.Errorf("scan search history: %w", err)
		}
		items = append(items, h)
	}
	return items, total, nil
}

func (r *SearchHistoryRepository) Delete(id int64) error {
	_, err := r.db.Exec("DELETE FROM search_history WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("delete search history: %w", err)
	}
	return nil
}

func (r *SearchHistoryRepository) Clear() error {
	_, err := r.db.Exec("DELETE FROM search_history")
	if err != nil {
		return fmt.Errorf("clear search history: %w", err)
	}
	return nil
}
