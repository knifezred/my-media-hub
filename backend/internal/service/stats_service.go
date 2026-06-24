package service

import (
	"database/sql"
	"my-media-hub/backend/internal/model"
	"my-media-hub/backend/internal/repository"
)

type StatsService struct {
	statsRepo *repository.StatsRepository
}

func NewStatsService(db *sql.DB) *StatsService {
	return &StatsService{
		statsRepo: repository.NewStatsRepository(db),
	}
}

func (s *StatsService) Overview() (*model.StatsOverview, error) {
	return s.statsRepo.Overview()
}
