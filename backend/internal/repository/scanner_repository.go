package repository

import (
	"database/sql"
	"fmt"
)

type ScannerRepository struct {
	db *sql.DB
}

func NewScannerRepository(db *sql.DB) *ScannerRepository {
	return &ScannerRepository{db: db}
}

func (r *ScannerRepository) ExistsByPath(path string) (bool, error) {
	var count int64
	err := r.db.QueryRow("SELECT COUNT(*) FROM media WHERE path = ?", path).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("check path exists: %w", err)
	}
	return count > 0, nil
}

func (r *ScannerRepository) ExistsByHash(hash string) (bool, error) {
	var count int64
	err := r.db.QueryRow("SELECT COUNT(*) FROM media WHERE hash = ?", hash).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("check hash exists: %w", err)
	}
	return count > 0, nil
}

func (r *ScannerRepository) BulkInsert(media []struct {
	MediaType string
	Title     string
	Path      string
	Hash      string
	Size      int64
}) ([]int64, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, fmt.Errorf("begin transaction: %w", err)
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(`
		INSERT INTO media (media_type, title, description, path, hash, size, cover_path)
		VALUES (?, '', ?, ?, ?, ?, '')`)
	if err != nil {
		return nil, fmt.Errorf("prepare insert: %w", err)
	}
	defer stmt.Close()

	var ids []int64
	for _, m := range media {
		result, err := stmt.Exec(m.MediaType, m.Title, m.Path, m.Hash, m.Size)
		if err != nil {
			return nil, fmt.Errorf("bulk insert media: %w", err)
		}
		id, err := result.LastInsertId()
		if err != nil {
			return nil, fmt.Errorf("get last insert id: %w", err)
		}
		ids = append(ids, id)
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("commit bulk insert: %w", err)
	}

	return ids, nil
}
