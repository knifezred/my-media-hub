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

func (r *BehaviorRepository) Insert(mediaID int64, behaviorType string, score float64) (int64, error) {
	result, err := r.db.Exec(
		"INSERT INTO user_behavior (media_id, behavior_type, score) VALUES (?, ?, ?)",
		mediaID, behaviorType, score,
	)
	if err != nil {
		return 0, fmt.Errorf("insert behavior: %w", err)
	}
	return result.LastInsertId()
}

func (r *BehaviorRepository) GetByMediaID(mediaID int64) ([]model.UserBehavior, error) {
	rows, err := r.db.Query(
		"SELECT id, media_id, behavior_type, score, created_at FROM user_behavior WHERE media_id = ? ORDER BY created_at DESC",
		mediaID,
	)
	if err != nil {
		return nil, fmt.Errorf("get behaviors by media: %w", err)
	}
	defer rows.Close()

	items := make([]model.UserBehavior, 0)
	for rows.Next() {
		var b model.UserBehavior
		if err := rows.Scan(&b.ID, &b.MediaID, &b.BehaviorType, &b.Score, &b.CreatedAt); err != nil {
			return nil, fmt.Errorf("scan behavior: %w", err)
		}
		items = append(items, b)
	}
	return items, rows.Err()
}

func (r *BehaviorRepository) GetBehaviorCount(mediaID int64, behaviorType string) (int64, error) {
	var count int64
	err := r.db.QueryRow(
		"SELECT COUNT(*) FROM user_behavior WHERE media_id = ? AND behavior_type = ?",
		mediaID, behaviorType,
	).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("count behavior: %w", err)
	}
	return count, nil
}

func (r *BehaviorRepository) GetTotalCount(behaviorType string) (int64, error) {
	var count int64
	query := "SELECT COUNT(*) FROM user_behavior"
	args := []interface{}{}
	if behaviorType != "" {
		query += " WHERE behavior_type = ?"
		args = append(args, behaviorType)
	}
	err := r.db.QueryRow(query, args...).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("count total behavior: %w", err)
	}
	return count, nil
}

func (r *BehaviorRepository) GetDistinctCount(behaviorType string) (int64, error) {
	var count int64
	query := "SELECT COUNT(DISTINCT media_id) FROM user_behavior WHERE behavior_type = ?"
	err := r.db.QueryRow(query, behaviorType).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("count distinct behavior: %w", err)
	}
	return count, nil
}

func (r *BehaviorRepository) GetAvgRating(mediaID int64) (float64, error) {
	var avg sql.NullFloat64
	err := r.db.QueryRow(
		"SELECT AVG(score) FROM user_behavior WHERE media_id = ? AND behavior_type = 'rating'",
		mediaID,
	).Scan(&avg)
	if err != nil {
		return 0, fmt.Errorf("get avg rating: %w", err)
	}
	return avg.Float64, nil
}

func (r *BehaviorRepository) GetMaxScoreByType(mediaID int64, behaviorType string) (float64, error) {
	var score sql.NullFloat64
	err := r.db.QueryRow(
		"SELECT MAX(score) FROM user_behavior WHERE media_id = ? AND behavior_type = ?",
		mediaID, behaviorType,
	).Scan(&score)
	if err != nil {
		return 0, fmt.Errorf("get max score: %w", err)
	}
	return score.Float64, nil
}

func (r *BehaviorRepository) DeleteByMediaIDAndType(mediaID int64, behaviorType string) error {
	_, err := r.db.Exec(
		"DELETE FROM user_behavior WHERE media_id = ? AND behavior_type = ?",
		mediaID, behaviorType,
	)
	if err != nil {
		return fmt.Errorf("delete behavior: %w", err)
	}
	return nil
}

func (r *BehaviorRepository) ListByType(behaviorType string, page, pageSize int) ([]model.UserBehavior, int64, error) {
	var total int64
	err := r.db.QueryRow(
		"SELECT COUNT(*) FROM user_behavior WHERE behavior_type = ?", behaviorType,
	).Scan(&total)
	if err != nil {
		return nil, 0, fmt.Errorf("count behaviors: %w", err)
	}

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}
	offset := (page - 1) * pageSize

	rows, err := r.db.Query(
		`SELECT id, media_id, behavior_type, score, created_at FROM user_behavior
		 WHERE behavior_type = ? ORDER BY created_at DESC LIMIT ? OFFSET ?`,
		behaviorType, pageSize, offset,
	)
	if err != nil {
		return nil, 0, fmt.Errorf("list behaviors: %w", err)
	}
	defer rows.Close()

	items := make([]model.UserBehavior, 0)
	for rows.Next() {
		var b model.UserBehavior
		if err := rows.Scan(&b.ID, &b.MediaID, &b.BehaviorType, &b.Score, &b.CreatedAt); err != nil {
			return nil, 0, fmt.Errorf("scan behavior: %w", err)
		}
		items = append(items, b)
	}
	return items, total, rows.Err()
}

func (r *BehaviorRepository) GetRecentByType(behaviorType string, limit int) ([]model.UserBehavior, error) {
	rows, err := r.db.Query(
		`SELECT id, media_id, behavior_type, score, created_at FROM user_behavior
		 WHERE behavior_type = ? ORDER BY created_at DESC LIMIT ?`,
		behaviorType, limit,
	)
	if err != nil {
		return nil, fmt.Errorf("get recent behaviors: %w", err)
	}
	defer rows.Close()

	items := make([]model.UserBehavior, 0)
	for rows.Next() {
		var b model.UserBehavior
		if err := rows.Scan(&b.ID, &b.MediaID, &b.BehaviorType, &b.Score, &b.CreatedAt); err != nil {
			return nil, fmt.Errorf("scan behavior: %w", err)
		}
		items = append(items, b)
	}
	return items, rows.Err()
}

func (r *BehaviorRepository) CountDistinctMediaByType(behaviorType string) (int64, error) {
	var count int64
	err := r.db.QueryRow(
		"SELECT COUNT(DISTINCT media_id) FROM user_behavior WHERE behavior_type = ?", behaviorType,
	).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("count distinct media by type: %w", err)
	}
	return count, nil
}

const mediaJoinFields = `m.id, m.media_type, m.title, m.description, m.path, m.hash, m.size, m.cover_path,
	m.favorite_count, m.view_count, m.rating_count, m.avg_rating, m.last_viewed_at, m.created_at, m.updated_at`

func (r *BehaviorRepository) ListMediaByType(behaviorType string, page, pageSize int) ([]model.Media, int64, error) {
	total, err := r.CountDistinctMediaByType(behaviorType)
	if err != nil {
		return nil, 0, err
	}

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}
	offset := (page - 1) * pageSize

	rows, err := r.db.Query(
		`SELECT `+mediaJoinFields+` FROM user_behavior ub
		 JOIN media m ON m.id = ub.media_id
		 WHERE ub.behavior_type = ?
		 GROUP BY ub.media_id
		 ORDER BY MAX(ub.created_at) DESC LIMIT ? OFFSET ?`,
		behaviorType, pageSize, offset,
	)
	if err != nil {
		return nil, 0, fmt.Errorf("list media by behavior: %w", err)
	}
	defer rows.Close()

	items := make([]model.Media, 0)
	for rows.Next() {
		var m model.Media
		var lastViewedAt sql.NullTime
		err := rows.Scan(
			&m.ID, &m.MediaType, &m.Title, &m.Description, &m.Path, &m.Hash, &m.Size, &m.CoverPath,
			&m.FavoriteCount, &m.ViewCount, &m.RatingCount, &m.AvgRating, &lastViewedAt, &m.CreatedAt, &m.UpdatedAt,
		)
		if err != nil {
			return nil, 0, fmt.Errorf("scan media from behavior join: %w", err)
		}
		if lastViewedAt.Valid {
			m.LastViewedAt = &lastViewedAt.Time
		}
		items = append(items, m)
	}
	return items, total, rows.Err()
}

func (r *BehaviorRepository) GetBehaviorStatistics() (*model.BehaviorStatistics, error) {
	stats := &model.BehaviorStatistics{}
	err := r.db.QueryRow("SELECT COUNT(*) FROM user_behavior WHERE behavior_type = 'favorite'").Scan(&stats.FavoriteCount)
	if err != nil {
		return nil, fmt.Errorf("count favorites: %w", err)
	}
	err = r.db.QueryRow("SELECT COUNT(*) FROM user_behavior WHERE behavior_type = 'view'").Scan(&stats.ViewCount)
	if err != nil {
		return nil, fmt.Errorf("count views: %w", err)
	}
	err = r.db.QueryRow("SELECT COUNT(*) FROM user_behavior WHERE behavior_type = 'rating'").Scan(&stats.RatingCount)
	if err != nil {
		return nil, fmt.Errorf("count ratings: %w", err)
	}
	err = r.db.QueryRow("SELECT COUNT(*) FROM user_behavior WHERE behavior_type = 'hidden'").Scan(&stats.HiddenCount)
	if err != nil {
		return nil, fmt.Errorf("count hidden: %w", err)
	}
	return stats, nil
}
