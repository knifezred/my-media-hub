package repository

import (
	"database/sql"
	"fmt"
)

type MediaCategoryRepository struct {
	db *sql.DB
}

func NewMediaCategoryRepository(db *sql.DB) *MediaCategoryRepository {
	return &MediaCategoryRepository{db: db}
}

func (r *MediaCategoryRepository) Add(mediaID, categoryID int64) error {
	_, err := r.db.Exec("INSERT OR IGNORE INTO media_category (media_id, category_id) VALUES (?, ?)", mediaID, categoryID)
	if err != nil {
		return fmt.Errorf("add media category: %w", err)
	}
	return nil
}

func (r *MediaCategoryRepository) Remove(mediaID, categoryID int64) error {
	_, err := r.db.Exec("DELETE FROM media_category WHERE media_id = ? AND category_id = ?", mediaID, categoryID)
	if err != nil {
		return fmt.Errorf("remove media category: %w", err)
	}
	return nil
}

func (r *MediaCategoryRepository) RemoveByMedia(mediaID int64) error {
	_, err := r.db.Exec("DELETE FROM media_category WHERE media_id = ?", mediaID)
	if err != nil {
		return fmt.Errorf("remove media categories: %w", err)
	}
	return nil
}
