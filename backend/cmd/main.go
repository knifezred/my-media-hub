package main

import (
	"log"

	"my-media-hub/backend/internal/config"
	"my-media-hub/backend/internal/database"
	"my-media-hub/backend/internal/router"
)

func main() {
	cfg := config.Load()

	db, err := database.Init(cfg.DatabasePath)
	if err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}

	r := router.Setup(db)

	log.Printf("server starting on %s", cfg.Addr)
	if err := r.Run(cfg.Addr); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
