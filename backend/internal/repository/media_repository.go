package repository

import (
	"database/sql"
	"fmt"
	"my-media-hub/backend/internal/model"
)

type MediaRepository struct {
	db *sql.DB
}

func NewMediaRepository(db *sql.DB) *MediaRepository {
	return &MediaRepository{db: db}
}

const mediaFields = `id, media_type, title, description, path, hash, size, cover_path,
	status, last_error, metadata_json, metadata_version,
	favorite, favorite_at, rating, rating_at, hidden, hidden_at,
	view_count, last_viewed_at, created_at, updated_at`

func scanMedia(s scanner) (*model.Media, error) {
	m := &model.Media{}
	var favoriteAt, ratingAt, hiddenAt, lastViewedAt sql.NullString
	err := s.Scan(
		&m.ID, &m.MediaType, &m.Title, &m.Description, &m.Path, &m.Hash, &m.Size, &m.CoverPath,
		&m.Status, &m.LastError, &m.MetadataJSON, &m.MetadataVersion,
		&m.Favorite, &favoriteAt, &m.Rating, &ratingAt, &m.Hidden, &hiddenAt,
		&m.ViewCount, &lastViewedAt, &m.CreatedAt, &m.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("scan media: %w", err)
	}
	m.FavoriteAt = parseTime(favoriteAt)
	m.RatingAt = parseTime(ratingAt)
	m.HiddenAt = parseTime(hiddenAt)
	m.LastViewedAt = parseTime(lastViewedAt)
	return m, nil
}

type scanner interface {
	Scan(dest ...interface{}) error
}

func (r *MediaRepository) GetByID(id int64) (*model.Media, error) {
	row := r.db.QueryRow(`SELECT `+mediaFields+` FROM media WHERE id = ?`, id)
	return scanMedia(row)
}

func (r *MediaRepository) GetByHash(hash string) (*model.Media, error) {
	row := r.db.QueryRow(`SELECT `+mediaFields+` FROM media WHERE hash = ?`, hash)
	return scanMedia(row)
}

func (r *MediaRepository) GetByPath(path string) (*model.Media, error) {
	row := r.db.QueryRow(`SELECT `+mediaFields+` FROM media WHERE path = ?`, path)
	return scanMedia(row)
}

func (r *MediaRepository) List(req model.MediaPageRequest) ([]model.Media, int64, error) {
	where := "WHERE 1=1"
	args := []interface{}{}

	if req.MediaType != "" {
		where += " AND media_type = ?"
		args = append(args, req.MediaType)
	}
	if req.Status != "" {
		where += " AND status = ?"
		args = append(args, req.Status)
	}

	var total int64
	if err := r.db.QueryRow(`SELECT COUNT(*) FROM media `+where, args...).Scan(&total); err != nil {
		return nil, 0, fmt.Errorf("count media: %w", err)
	}

	page := req.Page
	if page < 1 {
		page = 1
	}
	pageSize := req.PageSize
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}
	offset := (page - 1) * pageSize

	order := "ORDER BY created_at DESC"
	switch req.Sort {
	case "title":
		order = "ORDER BY title ASC"
	case "size":
		order = "ORDER BY size DESC"
	case "random":
		order = "ORDER BY RANDOM()"
	case "viewed":
		order = "ORDER BY last_viewed_at DESC"
	}

	rows, err := r.db.Query(
		`SELECT `+mediaFields+` FROM media `+where+` `+order+` LIMIT ? OFFSET ?`,
		append(args, pageSize, offset)...,
	)
	if err != nil {
		return nil, 0, fmt.Errorf("list media: %w", err)
	}
	defer rows.Close()

	items := make([]model.Media, 0)
	for rows.Next() {
		m, err := scanMedia(rows)
		if err != nil {
			return nil, 0, err
		}
		items = append(items, *m)
	}
	return items, total, nil
}

func (r *MediaRepository) Insert(m *model.Media) (int64, error) {
	res, err := r.db.Exec(`
		INSERT INTO media (media_type, title, description, path, hash, size, cover_path, status, metadata_json)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		m.MediaType, m.Title, m.Description, m.Path, m.Hash, m.Size, m.CoverPath,
		m.Status, m.MetadataJSON,
	)
	if err != nil {
		return 0, fmt.Errorf("insert media: %w", err)
	}
	return res.LastInsertId()
}

func (r *MediaRepository) Update(m *model.Media) error {
	_, err := r.db.Exec(`
		UPDATE media SET media_type=?, title=?, description=?, path=?, hash=?, size=?, cover_path=?,
			metadata_json=?, metadata_version=?
		WHERE id=?`,
		m.MediaType, m.Title, m.Description, m.Path, m.Hash, m.Size, m.CoverPath,
		m.MetadataJSON, m.MetadataVersion, m.ID,
	)
	return err
}

func (r *MediaRepository) UpdateStatus(id int64, status, lastError string) error {
	_, err := r.db.Exec(`UPDATE media SET status=?, last_error=? WHERE id=?`,
		status, lastError, id)
	return err
}

// 状态字段直接更新（favorite / rating / hidden / view）
func (r *MediaRepository) SetFavorite(id int64, favorite bool) error {
	if favorite {
		_, err := r.db.Exec(`UPDATE media SET favorite=1, favorite_at=strftime('%Y-%m-%dT%H:%M:%fZ','now') WHERE id=?`, id)
		return err
	}
	_, err := r.db.Exec(`UPDATE media SET favorite=0, favorite_at=NULL WHERE id=?`, id)
	return err
}

func (r *MediaRepository) SetRating(id int64, rating float64) error {
	if rating <= 0 {
		_, err := r.db.Exec(`UPDATE media SET rating=0, rating_at=NULL WHERE id=?`, id)
		return err
	}
	_, err := r.db.Exec(`UPDATE media SET rating=?, rating_at=strftime('%Y-%m-%dT%H:%M:%fZ','now') WHERE id=?`, rating, id)
	return err
}

func (r *MediaRepository) SetHidden(id int64, hidden bool) error {
	if hidden {
		_, err := r.db.Exec(`UPDATE media SET hidden=1, hidden_at=strftime('%Y-%m-%dT%H:%M:%fZ','now') WHERE id=?`, id)
		return err
	}
	_, err := r.db.Exec(`UPDATE media SET hidden=0, hidden_at=NULL WHERE id=?`, id)
	return err
}

func (r *MediaRepository) IncViewCount(id int64) error {
	_, err := r.db.Exec(`
		UPDATE media SET view_count=view_count+1, last_viewed_at=strftime('%Y-%m-%dT%H:%M:%fZ','now')
		WHERE id=?`, id)
	return err
}

func (r *MediaRepository) Delete(id int64) error {
	_, err := r.db.Exec(`DELETE FROM media WHERE id=?`, id)
	return err
}

func (r *MediaRepository) Count() (int64, error) {
	var count int64
	err := r.db.QueryRow(`SELECT COUNT(*) FROM media`).Scan(&count)
	return count, err
}

func (r *MediaRepository) CountByType() (images, videos, novels, music int64, err error) {
	rows, err := r.db.Query(`SELECT media_type, COUNT(*) FROM media GROUP BY media_type`)
	if err != nil {
		return 0, 0, 0, 0, err
	}
	defer rows.Close()
	for rows.Next() {
		var t string
		var c int64
		if err := rows.Scan(&t, &c); err != nil {
			return 0, 0, 0, 0, err
		}
		switch t {
		case model.MediaTypeImage:
			images = c
		case model.MediaTypeVideo:
			videos = c
		case model.MediaTypeNovel:
			novels = c
		case model.MediaTypeMusic:
			music = c
		}
	}
	return
}
