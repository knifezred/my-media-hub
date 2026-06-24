package repository

import (
	"database/sql"
	"fmt"
	"my-media-hub/backend/internal/model"
)

type TagRepository struct {
	db *sql.DB
}

func NewTagRepository(db *sql.DB) *TagRepository {
	return &TagRepository{db: db}
}

// Create 标签去重创建（基于 name_norm）
func (r *TagRepository) Create(name, source string) (*model.Tag, error) {
	nameNorm := normalizeTagName(name)
	res, err := r.db.Exec(`INSERT OR IGNORE INTO tag (name, name_norm, source) VALUES (?, ?, ?)`,
		name, nameNorm, source)
	if err != nil {
		return nil, fmt.Errorf("create tag: %w", err)
	}
	id, _ := res.LastInsertId()
	if id == 0 {
		// 已存在，通过 name_norm 查找
		return r.GetByNameNorm(nameNorm)
	}
	return &model.Tag{ID: id, Name: name, NameNorm: nameNorm, Source: source}, nil
}

func (r *TagRepository) GetByID(id int64) (*model.Tag, error) {
	row := r.db.QueryRow(`SELECT id, name, name_norm, source, created_at FROM tag WHERE id = ?`, id)
	t := &model.Tag{}
	err := row.Scan(&t.ID, &t.Name, &t.NameNorm, &t.Source, &t.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("get tag: %w", err)
	}
	return t, nil
}

func (r *TagRepository) GetByName(name string) (*model.Tag, error) {
	return r.GetByNameNorm(normalizeTagName(name))
}

func (r *TagRepository) GetByNameNorm(nameNorm string) (*model.Tag, error) {
	row := r.db.QueryRow(`SELECT id, name, name_norm, source, created_at FROM tag WHERE name_norm = ?`, nameNorm)
	t := &model.Tag{}
	err := row.Scan(&t.ID, &t.Name, &t.NameNorm, &t.Source, &t.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("get tag by name_norm: %w", err)
	}
	return t, nil
}

func (r *TagRepository) List(page, pageSize int) ([]model.Tag, int64, error) {
	var total int64
	if err := r.db.QueryRow(`SELECT COUNT(*) FROM tag`).Scan(&total); err != nil {
		return nil, 0, err
	}
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}
	offset := (page - 1) * pageSize
	rows, err := r.db.Query(`SELECT id, name, name_norm, source, created_at FROM tag ORDER BY id ASC LIMIT ? OFFSET ?`,
		pageSize, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("list tags: %w", err)
	}
	defer rows.Close()

	items := make([]model.Tag, 0)
	for rows.Next() {
		var t model.Tag
		if err := rows.Scan(&t.ID, &t.Name, &t.NameNorm, &t.Source, &t.CreatedAt); err != nil {
			return nil, 0, err
		}
		items = append(items, t)
	}
	return items, total, nil
}

func (r *TagRepository) GetByMediaID(mediaID int64) ([]model.Tag, error) {
	rows, err := r.db.Query(`
		SELECT t.id, t.name, t.name_norm, t.source, t.created_at
		FROM tag t JOIN media_tag mt ON mt.tag_id = t.id
		WHERE mt.media_id = ?`, mediaID)
	if err != nil {
		return nil, fmt.Errorf("get tags by media: %w", err)
	}
	defer rows.Close()

	items := make([]model.Tag, 0)
	for rows.Next() {
		var t model.Tag
		if err := rows.Scan(&t.ID, &t.Name, &t.NameNorm, &t.Source, &t.CreatedAt); err != nil {
			return nil, err
		}
		items = append(items, t)
	}
	return items, nil
}
