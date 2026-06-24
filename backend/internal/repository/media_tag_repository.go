package repository

import (
	"database/sql"
	"fmt"
)

type MediaTagRepository struct {
	db *sql.DB
}

func NewMediaTagRepository(db *sql.DB) *MediaTagRepository {
	return &MediaTagRepository{db: db}
}

func (r *MediaTagRepository) Add(mediaID, tagID int64) error {
	_, err := r.db.Exec("INSERT OR IGNORE INTO media_tag (media_id, tag_id) VALUES (?, ?)", mediaID, tagID)
	if err != nil {
		return fmt.Errorf("add media tag: %w", err)
	}
	return nil
}

func (r *MediaTagRepository) Remove(mediaID, tagID int64) error {
	_, err := r.db.Exec("DELETE FROM media_tag WHERE media_id = ? AND tag_id = ?", mediaID, tagID)
	if err != nil {
		return fmt.Errorf("remove media tag: %w", err)
	}
	return nil
}

func (r *MediaTagRepository) RemoveByMedia(mediaID int64) error {
	_, err := r.db.Exec("DELETE FROM media_tag WHERE media_id = ?", mediaID)
	if err != nil {
		return fmt.Errorf("remove media tags: %w", err)
	}
	return nil
}
