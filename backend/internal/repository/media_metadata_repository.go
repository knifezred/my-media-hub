package repository

import (
	"database/sql"
	"fmt"
)

type MediaMetadataRepository struct {
	db *sql.DB
}

func NewMediaMetadataRepository(db *sql.DB) *MediaMetadataRepository {
	return &MediaMetadataRepository{db: db}
}

func (r *MediaMetadataRepository) Upsert(mediaID int64, metaKey, metaValue string) error {
	_, err := r.db.Exec(`
		INSERT INTO media_metadata (media_id, meta_key, meta_value) VALUES (?, ?, ?)
		ON CONFLICT DO NOTHING`,
		mediaID, metaKey, metaValue)
	if err != nil {
		return fmt.Errorf("upsert media metadata: %w", err)
	}
	return nil
}

func (r *MediaMetadataRepository) GetByMediaID(mediaID int64) (map[string]string, error) {
	rows, err := r.db.Query("SELECT meta_key, meta_value FROM media_metadata WHERE media_id = ?", mediaID)
	if err != nil {
		return nil, fmt.Errorf("get media metadata: %w", err)
	}
	defer rows.Close()

	result := make(map[string]string)
	for rows.Next() {
		var metaKey, metaValue string
		if err := rows.Scan(&metaKey, &metaValue); err != nil {
			return nil, fmt.Errorf("scan media metadata: %w", err)
		}
		result[metaKey] = metaValue
	}
	return result, nil
}

func (r *MediaMetadataRepository) DeleteByMedia(mediaID int64) error {
	_, err := r.db.Exec("DELETE FROM media_metadata WHERE media_id = ?", mediaID)
	if err != nil {
		return fmt.Errorf("delete media metadata: %w", err)
	}
	return nil
}
