package repository

import (
	"database/sql"
	"fmt"
	"my-media-hub/backend/internal/model"
)

type BehaviorRepository struct {
	db *sql.DB
}

func NewBehaviorRepository(db *sql.DB) *BehaviorRepository {
	return &BehaviorRepository{db: db}
}

func (r *BehaviorRepository) Insert(mediaID int64, behaviorType, behaviorValue, behaviorSource string) (int64, error) {
	res, err := r.db.Exec(`
		INSERT INTO media_behavior (media_id, behavior_type, behavior_value, behavior_source)
		VALUES (?, ?, ?, ?)`,
		mediaID, behaviorType, behaviorValue, behaviorSource)
	if err != nil {
		return 0, fmt.Errorf("insert behavior: %w", err)
	}
	return res.LastInsertId()
}

func (r *BehaviorRepository) ListByType(behaviorType string, page, pageSize int) ([]model.MediaBehavior, int64, error) {
	var total int64
	if err := r.db.QueryRow(`SELECT COUNT(*) FROM media_behavior WHERE behavior_type = ?`, behaviorType).Scan(&total); err != nil {
		return nil, 0, err
	}
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}
	offset := (page - 1) * pageSize
	rows, err := r.db.Query(`
		SELECT id, media_id, behavior_type, behavior_value, behavior_source, created_at
		FROM media_behavior WHERE behavior_type = ?
		ORDER BY created_at DESC LIMIT ? OFFSET ?`,
		behaviorType, pageSize, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("list behaviors: %w", err)
	}
	defer rows.Close()

	items := make([]model.MediaBehavior, 0)
	for rows.Next() {
		var b model.MediaBehavior
		if err := rows.Scan(&b.ID, &b.MediaID, &b.BehaviorType, &b.BehaviorValue, &b.BehaviorSource, &b.CreatedAt); err != nil {
			return nil, 0, err
		}
		items = append(items, b)
	}
	return items, total, nil
}

func (r *BehaviorRepository) GetByMediaID(mediaID int64, limit int) ([]model.MediaBehavior, error) {
	rows, err := r.db.Query(`
		SELECT id, media_id, behavior_type, behavior_value, behavior_source, created_at
		FROM media_behavior WHERE media_id = ?
		ORDER BY created_at DESC LIMIT ?`, mediaID, limit)
	if err != nil {
		return nil, fmt.Errorf("get behaviors by media: %w", err)
	}
	defer rows.Close()

	items := make([]model.MediaBehavior, 0)
	for rows.Next() {
		var b model.MediaBehavior
		if err := rows.Scan(&b.ID, &b.MediaID, &b.BehaviorType, &b.BehaviorValue, &b.BehaviorSource, &b.CreatedAt); err != nil {
			return nil, err
		}
		items = append(items, b)
	}
	return items, nil
}

func (r *BehaviorRepository) GetRecent(limit int) ([]model.MediaBehavior, error) {
	rows, err := r.db.Query(`
		SELECT id, media_id, behavior_type, behavior_value, behavior_source, created_at
		FROM media_behavior ORDER BY created_at DESC LIMIT ?`, limit)
	if err != nil {
		return nil, fmt.Errorf("get recent behaviors: %w", err)
	}
	defer rows.Close()

	items := make([]model.MediaBehavior, 0)
	for rows.Next() {
		var b model.MediaBehavior
		if err := rows.Scan(&b.ID, &b.MediaID, &b.BehaviorType, &b.BehaviorValue, &b.BehaviorSource, &b.CreatedAt); err != nil {
			return nil, err
		}
		items = append(items, b)
	}
	return items, nil
}

func (r *BehaviorRepository) CountByType(behaviorType string) (int64, error) {
	var count int64
	err := r.db.QueryRow(`SELECT COUNT(*) FROM media_behavior WHERE behavior_type = ?`, behaviorType).Scan(&count)
	return count, err
}

func (r *BehaviorRepository) CountDistinctMedia(behaviorType string) (int64, error) {
	var count int64
	err := r.db.QueryRow(`SELECT COUNT(DISTINCT media_id) FROM media_behavior WHERE behavior_type = ?`, behaviorType).Scan(&count)
	return count, err
}

func (r *BehaviorRepository) Statistics() (*model.BehaviorStatistics, error) {
	stats := &model.BehaviorStatistics{}
	if err := r.db.QueryRow(`SELECT COUNT(*) FROM media_behavior WHERE behavior_type = ?`, model.BehaviorFavorite).Scan(&stats.FavoriteCount); err != nil {
		return nil, err
	}
	if err := r.db.QueryRow(`SELECT COUNT(*) FROM media_behavior WHERE behavior_type = ?`, model.BehaviorView).Scan(&stats.ViewCount); err != nil {
		return nil, err
	}
	if err := r.db.QueryRow(`SELECT COUNT(*) FROM media_behavior WHERE behavior_type = ?`, model.BehaviorRate).Scan(&stats.RateCount); err != nil {
		return nil, err
	}
	if err := r.db.QueryRow(`SELECT COUNT(*) FROM media_behavior WHERE behavior_type = ?`, model.BehaviorHide).Scan(&stats.HideCount); err != nil {
		return nil, err
	}
	return stats, nil
}
