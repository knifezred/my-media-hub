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

func (r *CategoryRepository) Create(name string, parentID int64) (*model.Category, error) {
	result, err := r.db.Exec("INSERT OR IGNORE INTO category (name, parent_id) VALUES (?, ?)", name, parentID)
	if err != nil {
		return nil, fmt.Errorf("create category: %w", err)
	}
	id, _ := result.LastInsertId()
	if id == 0 {
		return r.GetByName(name)
	}
	return &model.Category{ID: id, Name: name, ParentID: parentID}, nil
}

func (r *CategoryRepository) GetByName(name string) (*model.Category, error) {
	row := r.db.QueryRow("SELECT id, name, parent_id FROM category WHERE name = ?", name)
	c := &model.Category{}
	err := row.Scan(&c.ID, &c.Name, &c.ParentID)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("get category by name: %w", err)
	}
	return c, nil
}

func (r *CategoryRepository) GetByID(id int64) (*model.Category, error) {
	row := r.db.QueryRow("SELECT id, name, parent_id FROM category WHERE id = ?", id)
	c := &model.Category{}
	err := row.Scan(&c.ID, &c.Name, &c.ParentID)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("get category by id: %w", err)
	}
	return c, nil
}

func (r *CategoryRepository) List(page, pageSize int) ([]model.Category, int64, error) {
	var total int64
	err := r.db.QueryRow("SELECT COUNT(*) FROM category").Scan(&total)
	if err != nil {
		return nil, 0, fmt.Errorf("count categories: %w", err)
	}

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}
	offset := (page - 1) * pageSize

	rows, err := r.db.Query("SELECT id, name, parent_id FROM category ORDER BY id ASC LIMIT ? OFFSET ?", pageSize, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("list categories: %w", err)
	}
	defer rows.Close()

	items := make([]model.Category, 0)
	for rows.Next() {
		var c model.Category
		if err := rows.Scan(&c.ID, &c.Name, &c.ParentID); err != nil {
			return nil, 0, fmt.Errorf("scan category: %w", err)
		}
		items = append(items, c)
	}
	return items, total, nil
}

func (r *CategoryRepository) GetByMediaID(mediaID int64) ([]model.Category, error) {
	rows, err := r.db.Query(`
		SELECT c.id, c.name, c.parent_id FROM category c
		JOIN media_category mc ON mc.category_id = c.id
		WHERE mc.media_id = ?`, mediaID)
	if err != nil {
		return nil, fmt.Errorf("get categories by media id: %w", err)
	}
	defer rows.Close()

	items := make([]model.Category, 0)
	for rows.Next() {
		var c model.Category
		if err := rows.Scan(&c.ID, &c.Name, &c.ParentID); err != nil {
			return nil, fmt.Errorf("scan category: %w", err)
		}
		items = append(items, c)
	}
	return items, nil
}
