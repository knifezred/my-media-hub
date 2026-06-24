package search

import (
	"database/sql"
)

type Doc struct {
	ID          uint64
	Title       string
	Description string
}

type SearchResult struct {
	ID    uint64  `json:"id"`
	Title string  `json:"title"`
	Score float64 `json:"score"`
}

type BatchIndexItem struct {
	ID          uint64
	Title       string
	Description string
}

type Index struct {
	db       *sql.DB
	strategy Strategy
}

func NewIndex(db *sql.DB) *Index {
	return &Index{
		db:       db,
		strategy: &likeStrategy{},
	}
}

func (si *Index) SetStrategy(s Strategy) {
	si.strategy = s
}

func (si *Index) IndexMedia(id uint64, title, description string) error {
	return nil
}

func (si *Index) RemoveMedia(id uint64) error {
	return nil
}

func (si *Index) Search(keyword string, page, pageSize int) ([]SearchResult, int, error) {
	return si.strategy.Search(si.db, keyword, page, pageSize)
}

func (si *Index) Suggestions(prefix string, limit int) ([]string, error) {
	return si.strategy.Suggestions(si.db, prefix, limit)
}

func (si *Index) BatchIndex(items []BatchIndexItem) error {
	return nil
}

func (si *Index) Count() (uint64, error) {
	return 0, nil
}

func (si *Index) Close() error {
	return nil
}

func (si *Index) ReindexFromDB(db *sql.DB) error {
	return nil
}
