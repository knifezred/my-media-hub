package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

func Init(dbPath string) (*sql.DB, error) {
	dir := filepath.Dir(dbPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, fmt.Errorf("create database directory: %w", err)
	}

	db, err := sql.Open("sqlite3", dbPath+"?_journal_mode=WAL&_busy_timeout=5000")
	if err != nil {
		return nil, fmt.Errorf("open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("ping database: %w", err)
	}

	if err := migrate(db); err != nil {
		return nil, fmt.Errorf("migrate database: %w", err)
	}

	log.Printf("database initialized: %s", dbPath)
	return db, nil
}

func migrate(db *sql.DB) error {
	schema := `
	CREATE TABLE IF NOT EXISTS media (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		media_type TEXT NOT NULL DEFAULT '',
		title TEXT NOT NULL DEFAULT '',
		description TEXT NOT NULL DEFAULT '',
		path TEXT NOT NULL DEFAULT '',
		hash TEXT NOT NULL DEFAULT '',
		size INTEGER NOT NULL DEFAULT 0,
		cover_path TEXT NOT NULL DEFAULT '',
		created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
	);

	CREATE INDEX IF NOT EXISTS idx_media_type ON media(media_type);
	CREATE INDEX IF NOT EXISTS idx_media_created_at ON media(created_at);
	CREATE INDEX IF NOT EXISTS idx_media_hash ON media(hash);
	CREATE INDEX IF NOT EXISTS idx_media_path ON media(path);

	CREATE TABLE IF NOT EXISTS tag (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL DEFAULT ''
	);
	CREATE UNIQUE INDEX IF NOT EXISTS uniq_tag_name ON tag(name);

	CREATE TABLE IF NOT EXISTS media_tag (
		media_id INTEGER NOT NULL,
		tag_id INTEGER NOT NULL
	);
	CREATE INDEX IF NOT EXISTS idx_media_tag_media ON media_tag(media_id);
	CREATE INDEX IF NOT EXISTS idx_media_tag_tag ON media_tag(tag_id);

	CREATE TABLE IF NOT EXISTS category (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL DEFAULT '',
		parent_id INTEGER NOT NULL DEFAULT 0
	);

	CREATE TABLE IF NOT EXISTS media_category (
		media_id INTEGER NOT NULL,
		category_id INTEGER NOT NULL
	);
	CREATE INDEX IF NOT EXISTS idx_media_category_media ON media_category(media_id);
	CREATE INDEX IF NOT EXISTS idx_media_category_category ON media_category(category_id);

	CREATE TABLE IF NOT EXISTS media_metadata (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		media_id INTEGER NOT NULL,
		key TEXT NOT NULL DEFAULT '',
		value TEXT NOT NULL DEFAULT ''
	);
	CREATE INDEX IF NOT EXISTS idx_media_metadata_media ON media_metadata(media_id);
	CREATE INDEX IF NOT EXISTS idx_media_metadata_key ON media_metadata(key);

	CREATE TABLE IF NOT EXISTS user_favorite (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		media_id INTEGER NOT NULL,
		created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS user_rating (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		media_id INTEGER NOT NULL,
		rating INTEGER NOT NULL DEFAULT 0,
		created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS user_viewed (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		media_id INTEGER NOT NULL,
		viewed_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS user_hidden (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		media_id INTEGER NOT NULL,
		created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS search_history (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		keyword TEXT NOT NULL DEFAULT '',
		result_count INTEGER NOT NULL DEFAULT 0,
		created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS search_click_history (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		keyword TEXT NOT NULL DEFAULT '',
		media_id INTEGER NOT NULL,
		position INTEGER NOT NULL DEFAULT 0,
		created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS recommendation_cache (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		recommendation_type TEXT NOT NULL DEFAULT '',
		media_id INTEGER NOT NULL,
		score REAL NOT NULL DEFAULT 0,
		generated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
	);

	CREATE VIRTUAL TABLE IF NOT EXISTS media_fts USING fts5(
		title,
		description,
		tags,
		author,
		metadata,
		content='media',
		content_rowid='id'
	);

	CREATE TRIGGER IF NOT EXISTS media_ai AFTER INSERT ON media BEGIN
		INSERT INTO media_fts(rowid, title, description)
		VALUES (new.id, new.title, new.description);
	END;

	CREATE TRIGGER IF NOT EXISTS media_ad AFTER DELETE ON media BEGIN
		INSERT INTO media_fts(media_fts, rowid, title, description)
		VALUES ('delete', old.id, old.title, old.description);
	END;

	CREATE TRIGGER IF NOT EXISTS media_au AFTER UPDATE ON media BEGIN
		INSERT INTO media_fts(media_fts, rowid, title, description)
		VALUES ('delete', old.id, old.title, old.description);
		INSERT INTO media_fts(rowid, title, description)
		VALUES (new.id, new.title, new.description);
	END;
	`

	_, err := db.Exec(schema)
	if err != nil {
		return fmt.Errorf("execute schema: %w", err)
	}

	return nil
}
