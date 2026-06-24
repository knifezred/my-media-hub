package search

import "database/sql"

type Strategy interface {
	Search(db *sql.DB, keyword string, page, pageSize int) ([]SearchResult, int, error)
	Suggestions(db *sql.DB, prefix string, limit int) ([]string, error)
}
