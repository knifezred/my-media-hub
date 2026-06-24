package repository

import (
	"database/sql"
	"fmt"
	"my-media-hub/backend/internal/model"
	"my-media-hub/backend/internal/search"
	"strings"
)

type SearchRepository struct {
	db    *sql.DB
	index *search.Index
}

func NewSearchRepository(db *sql.DB, index *search.Index) *SearchRepository {
	return &SearchRepository{db: db, index: index}
}

func (r *SearchRepository) Search(keyword, mediaType string, page, pageSize int) ([]model.Media, int64, error) {
	results, total, err := r.index.Search(keyword, page, pageSize)
	if err != nil {
		return nil, 0, fmt.Errorf("search index: %w", err)
	}

	if len(results) == 0 {
		return []model.Media{}, int64(total), nil
	}

	ids := make([]int64, len(results))
	for i, r := range results {
		ids[i] = int64(r.ID)
	}

	placeholders := make([]string, len(ids))
	idArgs := make([]interface{}, len(ids))
	for i, id := range ids {
		placeholders[i] = "?"
		idArgs[i] = id
	}

	query := fmt.Sprintf(`
		SELECT id, media_type, title, description, path, hash, size, cover_path, created_at, updated_at
		FROM media
		WHERE id IN (%s)`, strings.Join(placeholders, ","))

	rows, err := r.db.Query(query, idArgs...)
	if err != nil {
		return nil, 0, fmt.Errorf("fetch search results: %w", err)
	}
	defer rows.Close()

	mediaMap := make(map[int64]model.Media, len(ids))
	for rows.Next() {
		var m model.Media
		if err := rows.Scan(&m.ID, &m.MediaType, &m.Title, &m.Description, &m.Path, &m.Hash, &m.Size, &m.CoverPath, &m.CreatedAt, &m.UpdatedAt); err != nil {
			return nil, 0, fmt.Errorf("scan media: %w", err)
		}
		mediaMap[m.ID] = m
	}

	items := make([]model.Media, 0, len(ids))
	for _, id := range ids {
		if m, ok := mediaMap[id]; ok {
			items = append(items, m)
		}
	}

	if mediaType != "" {
		filtered := make([]model.Media, 0, len(items))
		for _, m := range items {
			if m.MediaType == mediaType {
				filtered = append(filtered, m)
			}
		}
		items = filtered
	}

	return items, int64(total), nil
}

func (r *SearchRepository) Suggestions(prefix string, limit int) ([]string, error) {
	return r.index.Suggestions(prefix, limit)
}
