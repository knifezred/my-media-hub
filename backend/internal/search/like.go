package search

import (
	"database/sql"
	"fmt"
	"strings"
)

type likeStrategy struct{}

func (s *likeStrategy) Search(db *sql.DB, keyword string, page, pageSize int) ([]SearchResult, int, error) {
	keyword = strings.TrimSpace(keyword)
	if keyword == "" {
		return nil, 0, nil
	}

	like := "%" + keyword + "%"

	var total int
	err := db.QueryRow(`SELECT COUNT(*) FROM media WHERE title LIKE ? OR description LIKE ?`, like, like).Scan(&total)
	if err != nil {
		return nil, 0, fmt.Errorf("count search results: %w", err)
	}
	if total == 0 {
		return nil, 0, nil
	}

	offset := (page - 1) * pageSize
	if offset < 0 {
		offset = 0
	}

	rows, err := db.Query(
		`SELECT id, title FROM media WHERE title LIKE ? OR description LIKE ? ORDER BY id DESC LIMIT ? OFFSET ?`,
		like, like, pageSize, offset,
	)
	if err != nil {
		return nil, 0, fmt.Errorf("search media: %w", err)
	}
	defer rows.Close()

	items := make([]SearchResult, 0, pageSize)
	for rows.Next() {
		var item SearchResult
		if err := rows.Scan(&item.ID, &item.Title); err != nil {
			return nil, 0, fmt.Errorf("scan search result: %w", err)
		}
		items = append(items, item)
	}
	if err := rows.Err(); err != nil {
		return nil, 0, fmt.Errorf("iterate search results: %w", err)
	}

	return items, total, nil
}

func (s *likeStrategy) Suggestions(db *sql.DB, prefix string, limit int) ([]string, error) {
	prefix = strings.TrimSpace(prefix)
	if prefix == "" {
		return nil, nil
	}

	like := prefix + "%"

	rows, err := db.Query(
		`SELECT DISTINCT title FROM media WHERE title LIKE ? LIMIT ?`,
		like, limit,
	)
	if err != nil {
		return nil, fmt.Errorf("search suggestions: %w", err)
	}
	defer rows.Close()

	items := make([]string, 0, limit)
	for rows.Next() {
		var title string
		if err := rows.Scan(&title); err != nil {
			return nil, fmt.Errorf("scan suggestion: %w", err)
		}
		items = append(items, title)
	}
	return items, rows.Err()
}
