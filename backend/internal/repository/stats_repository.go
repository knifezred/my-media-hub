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

	if err := r.db.QueryRow(`SELECT COUNT(*) FROM media`).Scan(&stats.TotalMedia); err != nil {
		return nil, fmt.Errorf("count media: %w", err)
	}

	rows, err := r.db.Query(`SELECT media_type, COUNT(*) FROM media GROUP BY media_type`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var t string
		var c int64
		if err := rows.Scan(&t, &c); err != nil {
			return nil, err
		}
		switch t {
		case model.MediaTypeImage:
			stats.TotalImages = c
		case model.MediaTypeVideo:
			stats.TotalVideos = c
		case model.MediaTypeNovel:
			stats.TotalNovels = c
		case model.MediaTypeMusic:
			stats.TotalMusic = c
		}
	}

	return stats, nil
}
