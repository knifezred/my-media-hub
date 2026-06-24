package repository

import (
	"database/sql"
	"fmt"
	"my-media-hub/backend/internal/model"
)

type StatsRepository struct {
	db *sql.DB
}

func NewStatsRepository(db *sql.DB) *StatsRepository {
	return &StatsRepository{db: db}
}

func (r *StatsRepository) Overview() (*model.StatsOverview, error) {
	stats := &model.StatsOverview{}

	err := r.db.QueryRow("SELECT COUNT(*) FROM media").Scan(&stats.TotalMedia)
	if err != nil {
		return nil, fmt.Errorf("count total media: %w", err)
	}

	rows, err := r.db.Query("SELECT media_type, COUNT(*) FROM media GROUP BY media_type")
	if err != nil {
		return nil, fmt.Errorf("count by type: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var mediaType string
		var count int64
		if err := rows.Scan(&mediaType, &count); err != nil {
			return nil, fmt.Errorf("scan count by type: %w", err)
		}
		switch mediaType {
		case "image":
			stats.TotalImages = count
		case "video":
			stats.TotalVideos = count
		case "novel":
			stats.TotalNovels = count
		}
	}

	err = r.db.QueryRow("SELECT COUNT(*) FROM user_behavior WHERE behavior_type = 'favorite'").Scan(&stats.FavoriteCount)
	if err != nil {
		return nil, fmt.Errorf("count favorites: %w", err)
	}

	err = r.db.QueryRow("SELECT COUNT(DISTINCT media_id) FROM user_behavior WHERE behavior_type = 'view'").Scan(&stats.ViewedCount)
	if err != nil {
		return nil, fmt.Errorf("count viewed: %w", err)
	}

	return stats, nil
}
