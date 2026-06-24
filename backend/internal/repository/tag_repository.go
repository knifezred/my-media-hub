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

func (r *TagRepository) Create(name string) (*model.Tag, error) {
	result, err := r.db.Exec("INSERT OR IGNORE INTO tag (name) VALUES (?)", name)
	if err != nil {
		return nil, fmt.Errorf("create tag: %w", err)
	}
	id, _ := result.LastInsertId()
	if id == 0 {
		return r.GetByName(name)
	}
	return &model.Tag{ID: id, Name: name}, nil
}

func (r *TagRepository) GetByName(name string) (*model.Tag, error) {
	row := r.db.QueryRow("SELECT id, name FROM tag WHERE name = ?", name)
	t := &model.Tag{}
	err := row.Scan(&t.ID, &t.Name)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("get tag by name: %w", err)
	}
	return t, nil
}

func (r *TagRepository) GetByID(id int64) (*model.Tag, error) {
	row := r.db.QueryRow("SELECT id, name FROM tag WHERE id = ?", id)
	t := &model.Tag{}
	err := row.Scan(&t.ID, &t.Name)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("get tag by id: %w", err)
	}
	return t, nil
}

func (r *TagRepository) List(page, pageSize int) ([]model.Tag, int64, error) {
	var total int64
	err := r.db.QueryRow("SELECT COUNT(*) FROM tag").Scan(&total)
	if err != nil {
		return nil, 0, fmt.Errorf("count tags: %w", err)
	}

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}
	offset := (page - 1) * pageSize

	rows, err := r.db.Query("SELECT id, name FROM tag ORDER BY id ASC LIMIT ? OFFSET ?", pageSize, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("list tags: %w", err)
	}
	defer rows.Close()

	items := make([]model.Tag, 0)
	for rows.Next() {
		var t model.Tag
		if err := rows.Scan(&t.ID, &t.Name); err != nil {
			return nil, 0, fmt.Errorf("scan tag: %w", err)
		}
		items = append(items, t)
	}
	return items, total, nil
}

func (r *TagRepository) GetByMediaID(mediaID int64) ([]model.Tag, error) {
	rows, err := r.db.Query(`
		SELECT t.id, t.name FROM tag t
		JOIN media_tag mt ON mt.tag_id = t.id
		WHERE mt.media_id = ?`, mediaID)
	if err != nil {
		return nil, fmt.Errorf("get tags by media id: %w", err)
	}
	defer rows.Close()

	items := make([]model.Tag, 0)
	for rows.Next() {
		var t model.Tag
		if err := rows.Scan(&t.ID, &t.Name); err != nil {
			return nil, fmt.Errorf("scan tag: %w", err)
		}
		items = append(items, t)
	}
	return items, nil
}
