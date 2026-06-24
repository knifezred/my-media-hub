package repository

import (
	"database/sql"
	"fmt"
	"my-media-hub/backend/internal/model"
)

type ScannerRepository struct {
	db *sql.DB
}

func NewScannerRepository(db *sql.DB) *ScannerRepository {
	return &ScannerRepository{db: db}
}

// ExistsByHash 检查 media 表是否已存在同 hash 资源
func (r *ScannerRepository) ExistsByHash(hash string) (bool, error) {
	var count int64
	err := r.db.QueryRow(`SELECT COUNT(*) FROM media WHERE hash = ?`, hash).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("check hash: %w", err)
	}
	return count > 0, nil
}

// ExistsByPath 检查 scanner_index 是否已记录该路径
func (r *ScannerRepository) ExistsByPath(path string) (bool, error) {
	var count int64
	err := r.db.QueryRow(`SELECT COUNT(*) FROM scanner_index WHERE file_path = ?`, path).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("check path: %w", err)
	}
	return count > 0, nil
}

// Upsert 写入或更新 scanner_index
func (r *ScannerRepository) Upsert(path string, size int64, modTime, hash string) error {
	_, err := r.db.Exec(`
		INSERT INTO scanner_index (file_path, file_size, modified_time, file_hash)
		VALUES (?, ?, ?, ?)
		ON CONFLICT(file_path) DO UPDATE SET
			file_size=excluded.file_size,
			modified_time=excluded.modified_time,
			file_hash=excluded.file_hash,
			last_scan_at=strftime('%Y-%m-%dT%H:%M:%fZ','now')`,
		path, size, modTime, hash)
	return err
}

// GetByPath 按路径查 scanner_index
func (r *ScannerRepository) GetByPath(path string) (*model.ScannerIndex, error) {
	row := r.db.QueryRow(`
		SELECT id, media_id, file_path, file_size, modified_time, file_hash, last_scan_at
		FROM scanner_index WHERE file_path = ?`, path)
	si := &model.ScannerIndex{}
	var mediaID sql.NullInt64
	err := row.Scan(&si.ID, &mediaID, &si.FilePath, &si.FileSize, &si.ModifiedTime, &si.FileHash, &si.LastScanAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("get scanner_index: %w", err)
	}
	if mediaID.Valid {
		v := mediaID.Int64
		si.MediaID = &v
	}
	return si, nil
}

// LinkMedia 将 scanner_index 与 media 表关联
func (r *ScannerRepository) LinkMedia(path string, mediaID int64) error {
	_, err := r.db.Exec(`UPDATE scanner_index SET media_id = ? WHERE file_path = ?`, mediaID, path)
	return err
}

// GetUnlinked 获取未入库的扫描记录
func (r *ScannerRepository) GetUnlinked() ([]model.ScannerIndex, error) {
	rows, err := r.db.Query(`
		SELECT id, media_id, file_path, file_size, modified_time, file_hash, last_scan_at
		FROM scanner_index WHERE media_id IS NULL
		ORDER BY file_path`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := make([]model.ScannerIndex, 0)
	for rows.Next() {
		var si model.ScannerIndex
		var mediaID sql.NullInt64
		if err := rows.Scan(&si.ID, &mediaID, &si.FilePath, &si.FileSize, &si.ModifiedTime, &si.FileHash, &si.LastScanAt); err != nil {
			return nil, err
		}
		if mediaID.Valid {
			v := mediaID.Int64
			si.MediaID = &v
		}
		items = append(items, si)
	}
	return items, nil
}
