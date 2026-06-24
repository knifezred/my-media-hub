package api

import (
	"github.com/gin-gonic/gin"
	"my-media-hub/backend/internal/errorcode"
	"my-media-hub/backend/internal/response"
	"my-media-hub/backend/internal/service"
)

type StatsAPI struct {
	svc *service.StatsService
}

func NewStatsAPI(svc *service.StatsService) *StatsAPI {
	return &StatsAPI{svc: svc}
}

func (h *StatsAPI) Register(r *gin.RouterGroup) {
	r.GET("/stats/overview", h.Overview)
}

func (h *StatsAPI) Overview(c *gin.Context) {
	stats, err := h.svc.Overview()
	if err != nil {
		response.Error(c, errorcode.InternalError)
		return
	}

	response.Success(c, stats)
}
