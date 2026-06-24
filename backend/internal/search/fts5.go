package search

import (
	"database/sql"
	"fmt"
	"strings"
)

type fts5Strategy struct{}

func (s *fts5Strategy) Search(db *sql.DB, keyword string, page, pageSize int) ([]SearchResult, int, error) {
	keyword = strings.TrimSpace(keyword)
	if keyword == "" {
		return nil, 0, nil
	}

	query := strings.ReplaceAll(keyword, `"`, `""`)
	query = `"` + query + `"`

	var total int
	err := db.QueryRow(`SELECT COUNT(*) FROM media_fts WHERE media_fts MATCH ?`, query).Scan(&total)
	if err != nil {
		return nil, 0, fmt.Errorf("count fts5 results: %w", err)
	}
	if total == 0 {
		return nil, 0, nil
	}

	offset := (page - 1) * pageSize
	if offset < 0 {
		offset = 0
	}

	rows, err := db.Query(
		`SELECT m.id, m.title FROM media_fts f JOIN media m ON m.id = f.rowid WHERE media_fts MATCH ? ORDER BY rank LIMIT ? OFFSET ?`,
		query, pageSize, offset,
	)
	if err != nil {
		return nil, 0, fmt.Errorf("fts5 search: %w", err)
	}
	defer rows.Close()

	items := make([]SearchResult, 0, pageSize)
	for rows.Next() {
		var item SearchResult
		if err := rows.Scan(&item.ID, &item.Title); err != nil {
			return nil, 0, fmt.Errorf("scan fts5 result: %w", err)
		}
		items = append(items, item)
	}
	if err := rows.Err(); err != nil {
		return nil, 0, fmt.Errorf("iterate fts5 results: %w", err)
	}

	return items, total, nil
}

func (s *fts5Strategy) Suggestions(db *sql.DB, prefix string, limit int) ([]string, error) {
	prefix = strings.TrimSpace(prefix)
	if prefix == "" {
		return nil, nil
	}

	query := strings.ReplaceAll(prefix, `"`, `""`)
	query = `"` + query + `*"`

	rows, err := db.Query(
		`SELECT DISTINCT m.title FROM media_fts f JOIN media m ON m.id = f.rowid WHERE media_fts MATCH ? LIMIT ?`,
		query, limit,
	)
	if err != nil {
		return nil, fmt.Errorf("fts5 suggestions: %w", err)
	}
	defer rows.Close()

	items := make([]string, 0, limit)
	for rows.Next() {
		var title string
		if err := rows.Scan(&title); err != nil {
			return nil, fmt.Errorf("scan fts5 suggestion: %w", err)
		}
		items = append(items, title)
	}
	return items, rows.Err()
}
