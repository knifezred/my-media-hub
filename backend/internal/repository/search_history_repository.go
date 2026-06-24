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

// InsertOrUpdate 去重插入搜索历史（同关键词聚合）
func (r *SearchHistoryRepository) InsertOrUpdate(keyword string) (int64, error) {
	norm := normalizeKeyword(keyword)
	res, err := r.db.Exec(`
		INSERT INTO search_history (keyword, keyword_norm, use_count)
		VALUES (?, ?, 1)
		ON CONFLICT(keyword_norm) DO UPDATE SET
			use_count = use_count + 1,
			last_used_at = strftime('%Y-%m-%dT%H:%M:%fZ','now')`,
		keyword, norm)
	if err != nil {
		return 0, fmt.Errorf("insert search history: %w", err)
	}
	return res.LastInsertId()
}

func (r *SearchHistoryRepository) List(page, pageSize int) ([]model.SearchHistory, int64, error) {
	var total int64
	if err := r.db.QueryRow(`SELECT COUNT(*) FROM search_history`).Scan(&total); err != nil {
		return nil, 0, err
	}
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}
	offset := (page - 1) * pageSize
	rows, err := r.db.Query(`
		SELECT id, keyword, keyword_norm, use_count, last_used_at, created_at
		FROM search_history ORDER BY last_used_at DESC LIMIT ? OFFSET ?`,
		pageSize, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("list search history: %w", err)
	}
	defer rows.Close()

	items := make([]model.SearchHistory, 0)
	for rows.Next() {
		var h model.SearchHistory
		if err := rows.Scan(&h.ID, &h.Keyword, &h.KeywordNorm, &h.UseCount, &h.LastUsedAt, &h.CreatedAt); err != nil {
			return nil, 0, err
		}
		items = append(items, h)
	}
	return items, total, nil
}

func (r *SearchHistoryRepository) Delete(id int64) error {
	_, err := r.db.Exec(`DELETE FROM search_history WHERE id = ?`, id)
	return err
}

func (r *SearchHistoryRepository) Clear() error {
	_, err := r.db.Exec(`DELETE FROM search_history`)
	return err
}
