package repository

import "database/sql"

type MediaTagRepository struct {
	db *sql.DB
}

func NewMediaTagRepository(db *sql.DB) *MediaTagRepository {
	return &MediaTagRepository{db: db}
}

func (r *MediaTagRepository) Add(mediaID, tagID int64) error {
	_, err := r.db.Exec(`INSERT OR IGNORE INTO media_tag (media_id, tag_id) VALUES (?, ?)`,
		mediaID, tagID)
	return err
}

func (r *MediaTagRepository) Remove(mediaID, tagID int64) error {
	_, err := r.db.Exec(`DELETE FROM media_tag WHERE media_id = ? AND tag_id = ?`,
		mediaID, tagID)
	return err
}

func (r *MediaTagRepository) RemoveByMedia(mediaID int64) error {
	_, err := r.db.Exec(`DELETE FROM media_tag WHERE media_id = ?`, mediaID)
	return err
}

type MediaCategoryRepository struct {
	db *sql.DB
}

func NewMediaCategoryRepository(db *sql.DB) *MediaCategoryRepository {
	return &MediaCategoryRepository{db: db}
}

func (r *MediaCategoryRepository) Add(mediaID, categoryID int64, isPrimary bool) error {
	prim := 0
	if isPrimary {
		prim = 1
	}
	_, err := r.db.Exec(`INSERT OR IGNORE INTO media_category (media_id, category_id, is_primary) VALUES (?, ?, ?)`,
		mediaID, categoryID, prim)
	return err
}

func (r *MediaCategoryRepository) Remove(mediaID, categoryID int64) error {
	_, err := r.db.Exec(`DELETE FROM media_category WHERE media_id = ? AND category_id = ?`,
		mediaID, categoryID)
	return err
}

func (r *MediaCategoryRepository) RemoveByMedia(mediaID int64) error {
	_, err := r.db.Exec(`DELETE FROM media_category WHERE media_id = ?`, mediaID)
	return err
}

