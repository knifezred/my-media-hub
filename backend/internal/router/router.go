package router

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"my-media-hub/backend/internal/middleware"
)

func Setup(db *sql.DB) *gin.Engine {
	r := gin.New()

	r.Use(middleware.Recovery())
	r.Use(middleware.Logger())

	api := r.Group("/api/v1")

	_ = db
	_ = api

	return r
}
