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

func (r *MediaRepository) Insert(m *model.Media) (int64, error) {
	result, err := r.db.Exec(`
		INSERT INTO media (media_type, title, description, path, hash, size, cover_path)
		VALUES (?, ?, ?, ?, ?, ?, ?)`,
		m.MediaType, m.Title, m.Description, m.Path, m.Hash, m.Size, m.CoverPath,
	)
	if err != nil {
		return 0, fmt.Errorf("insert media: %w", err)
	}
	return result.LastInsertId()
}

const mediaFields = `id, media_type, title, description, path, hash, size, cover_path,
	favorite_count, view_count, rating_count, avg_rating, last_viewed_at, created_at, updated_at`

func scanMedia(scanner interface {
	Scan(dest ...interface{}) error
}) (*model.Media, error) {
	m := &model.Media{}
	var lastViewedAt sql.NullTime
	err := scanner.Scan(
		&m.ID, &m.MediaType, &m.Title, &m.Description, &m.Path, &m.Hash, &m.Size, &m.CoverPath,
		&m.FavoriteCount, &m.ViewCount, &m.RatingCount, &m.AvgRating, &lastViewedAt, &m.CreatedAt, &m.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("scan media: %w", err)
	}
	if lastViewedAt.Valid {
		m.LastViewedAt = &lastViewedAt.Time
	}
	return m, nil
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
		where += " AND m.media_type = ?"
		args = append(args, req.MediaType)
	}

	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM media m %s", where)
	var total int64
	err := r.db.QueryRow(countQuery, args...).Scan(&total)
	if err != nil {
		return nil, 0, fmt.Errorf("count media: %w", err)
	}

	if req.Page < 1 {
		req.Page = 1
	}
	pageSize := req.PageSize
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}
	offset := (req.Page - 1) * pageSize

	order := "ORDER BY m.created_at DESC"
	switch req.Sort {
	case "title":
		order = "ORDER BY m.title ASC"
	case "size":
		order = "ORDER BY m.size DESC"
	case "random":
		order = "ORDER BY RANDOM()"
	}

	query := fmt.Sprintf(`
		SELECT m.`+mediaFields+` FROM media m %s %s LIMIT ? OFFSET ?`, where, order)
	args = append(args, pageSize, offset)

	rows, err := r.db.Query(query, args...)
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

func (r *MediaRepository) Update(m *model.Media) error {
	_, err := r.db.Exec(`
		UPDATE media SET media_type=?, title=?, description=?, path=?, hash=?, size=?, cover_path=?
		WHERE id=?`,
		m.MediaType, m.Title, m.Description, m.Path, m.Hash, m.Size, m.CoverPath, m.ID,
	)
	if err != nil {
		return fmt.Errorf("update media: %w", err)
	}
	return nil
}

func (r *MediaRepository) UpdateStats(id int64, favoriteCount, viewCount, ratingCount int64, avgRating float64) error {
	_, err := r.db.Exec(`
		UPDATE media SET favorite_count=?, view_count=?, rating_count=?, avg_rating=?
		WHERE id=?`,
		favoriteCount, viewCount, ratingCount, avgRating, id,
	)
	if err != nil {
		return fmt.Errorf("update media stats: %w", err)
	}
	return nil
}

func (r *MediaRepository) UpdateViewStats(id int64, viewCount int64) error {
	_, err := r.db.Exec(`
		UPDATE media SET view_count=?, last_viewed_at=CURRENT_TIMESTAMP WHERE id=?`,
		viewCount, id,
	)
	if err != nil {
		return fmt.Errorf("update view stats: %w", err)
	}
	return nil
}

func (r *MediaRepository) Delete(id int64) error {
	_, err := r.db.Exec("DELETE FROM media WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("delete media: %w", err)
	}
	return nil
}

func (r *MediaRepository) Count() (int64, error) {
	var count int64
	err := r.db.QueryRow("SELECT COUNT(*) FROM media").Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("count media: %w", err)
	}
	return count, nil
}

func (r *MediaRepository) CountByType() (images, videos, novels int64, err error) {
	rows, err := r.db.Query("SELECT media_type, COUNT(*) FROM media GROUP BY media_type")
	if err != nil {
		return 0, 0, 0, fmt.Errorf("count by type: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var mediaType string
		var count int64
		if err := rows.Scan(&mediaType, &count); err != nil {
			return 0, 0, 0, fmt.Errorf("scan count by type: %w", err)
		}
		switch mediaType {
		case "image":
			images = count
		case "video":
			videos = count
		case "novel":
			novels = count
		}
	}
	return
}
