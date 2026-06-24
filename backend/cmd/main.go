package main

import (
	"log"

	"my-media-hub/backend/internal/config"
	"my-media-hub/backend/internal/database"
	"my-media-hub/backend/internal/router"
	"my-media-hub/backend/internal/search"
)

func main() {
	cfg := config.Load()

	db, err := database.Init(cfg.DatabasePath)
	if err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}

	idx := search.NewIndex(db)

	r := router.Setup(db, idx)

	log.Printf("server starting on %s", cfg.Addr)
	if err := r.Run(cfg.Addr); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
