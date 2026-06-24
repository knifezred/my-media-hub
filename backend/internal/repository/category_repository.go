package repository

import (
	"database/sql"
	"fmt"
	"my-media-hub/backend/internal/model"
)

type CategoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func scanCategory(row scanner) (*model.Category, error) {
	c := &model.Category{}
	var parentID sql.NullInt64
	err := row.Scan(&c.ID, &c.Name, &parentID, &c.Level, &c.Path, &c.Sort, &c.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("scan category: %w", err)
	}
	if parentID.Valid {
		v := parentID.Int64
		c.ParentID = &v
	}
	return c, nil
}

func (r *CategoryRepository) Create(name string, parentID *int64, level int, path string, sort int) (*model.Category, error) {
	res, err := r.db.Exec(`
		INSERT INTO category (name, parent_id, level, path, sort)
		VALUES (?, ?, ?, ?, ?)`,
		name, nullInt64(parentID), level, path, sort)
	if err != nil {
		return nil, fmt.Errorf("create category: %w", err)
	}
	id, _ := res.LastInsertId()
	return &model.Category{ID: id, Name: name, ParentID: parentID, Level: level, Path: path, Sort: sort}, nil
}

func nullInt64(p *int64) sql.NullInt64 {
	if p == nil {
		return sql.NullInt64{}
	}
	return sql.NullInt64{Int64: *p, Valid: true}
}

func (r *CategoryRepository) GetByID(id int64) (*model.Category, error) {
	row := r.db.QueryRow(`SELECT id, name, parent_id, level, path, sort, created_at FROM category WHERE id = ?`, id)
	return scanCategory(row)
}

func (r *CategoryRepository) List(page, pageSize int) ([]model.Category, int64, error) {
	var total int64
	if err := r.db.QueryRow(`SELECT COUNT(*) FROM category`).Scan(&total); err != nil {
		return nil, 0, err
	}
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}
	offset := (page - 1) * pageSize
	rows, err := r.db.Query(`SELECT id, name, parent_id, level, path, sort, created_at FROM category ORDER BY sort ASC, id ASC LIMIT ? OFFSET ?`,
		pageSize, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("list categories: %w", err)
	}
	defer rows.Close()

	items := make([]model.Category, 0)
	for rows.Next() {
		c, err := scanCategory(rows)
		if err != nil {
			return nil, 0, err
		}
		items = append(items, *c)
	}
	return items, total, nil
}

func (r *CategoryRepository) GetByMediaID(mediaID int64) ([]model.Category, error) {
	rows, err := r.db.Query(`
		SELECT c.id, c.name, c.parent_id, c.level, c.path, c.sort, c.created_at
		FROM category c JOIN media_category mc ON mc.category_id = c.id
		WHERE mc.media_id = ?`, mediaID)
	if err != nil {
		return nil, fmt.Errorf("get categories by media: %w", err)
	}
	defer rows.Close()

	items := make([]model.Category, 0)
	for rows.Next() {
		c, err := scanCategory(rows)
		if err != nil {
			return nil, err
		}
		items = append(items, *c)
	}
	return items, nil
}
