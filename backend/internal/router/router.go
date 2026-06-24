package router

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"my-media-hub/backend/internal/api"
	"my-media-hub/backend/internal/middleware"
	"my-media-hub/backend/internal/search"
	"my-media-hub/backend/internal/service"
)

func Setup(db *sql.DB, idx *search.Index) *gin.Engine {
	r := gin.New()

	r.Use(middleware.Recovery())
	r.Use(middleware.Logger())

	apiGroup := r.Group("/api/v1")

	behaviorSvc := service.NewBehaviorService(db)

	mediaSvc := service.NewMediaService(db)
	mediaAPI := api.NewMediaAPI(mediaSvc)
	mediaAPI.Register(apiGroup)

	searchSvc := service.NewSearchService(db, idx)
	searchAPI := api.NewSearchAPI(searchSvc)
	searchAPI.Register(apiGroup)

	behaviorAPI := api.NewBehaviorAPI(behaviorSvc)
	behaviorAPI.Register(apiGroup)

	statsSvc := service.NewStatsService(db)
	statsAPI := api.NewStatsAPI(statsSvc)
	statsAPI.Register(apiGroup)

	scannerSvc := service.NewScannerService(db, idx)
	scannerAPI := api.NewScannerAPI(scannerSvc)
	scannerAPI.Register(apiGroup)

	return r
}
