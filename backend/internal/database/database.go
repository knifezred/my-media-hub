package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

func Init(dbPath string) (*sql.DB, error) {
	dir := filepath.Dir(dbPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, fmt.Errorf("create database directory: %w", err)
	}

	if err := os.Remove(dbPath); err != nil && !os.IsNotExist(err) {
		log.Printf("warning: failed to remove old database: %v", err)
	}

	db, err := sql.Open("sqlite", dbPath+"?mode=rwc&_journal_mode=WAL&_busy_timeout=5000")
	if err != nil {
		return nil, fmt.Errorf("open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("ping database: %w", err)
	}

	if err := migrate(db); err != nil {
		return nil, fmt.Errorf("migrate database: %w", err)
	}

	log.Printf("database initialized (ERD v2.1): %s", dbPath)
	return db, nil
}

func migrate(db *sql.DB) error {
	schema := `
	-- ==================== 媒体主表 ====================
	CREATE TABLE IF NOT EXISTS media (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		media_type TEXT NOT NULL DEFAULT '',
		title TEXT NOT NULL DEFAULT '',
		description TEXT NOT NULL DEFAULT '',
		path TEXT NOT NULL DEFAULT '',
		hash TEXT NOT NULL DEFAULT '',
		size INTEGER NOT NULL DEFAULT 0,
		cover_path TEXT NOT NULL DEFAULT '',
		status TEXT NOT NULL DEFAULT 'new',
		last_error TEXT NOT NULL DEFAULT '',
		metadata_json TEXT NOT NULL DEFAULT '{}',
		metadata_version INTEGER NOT NULL DEFAULT 1,
		favorite INTEGER NOT NULL DEFAULT 0,
		favorite_at TEXT,
		rating REAL NOT NULL DEFAULT 0,
		rating_at TEXT,
		hidden INTEGER NOT NULL DEFAULT 0,
		hidden_at TEXT,
		view_count INTEGER NOT NULL DEFAULT 0,
		last_viewed_at TEXT,
		created_at TEXT NOT NULL DEFAULT (strftime('%Y-%m-%dT%H:%M:%fZ','now')),
		updated_at TEXT NOT NULL DEFAULT (strftime('%Y-%m-%dT%H:%M:%fZ','now'))
	);
	CREATE INDEX IF NOT EXISTS idx_media_type ON media(media_type);
	CREATE INDEX IF NOT EXISTS idx_media_created_at ON media(created_at);
	CREATE INDEX IF NOT EXISTS idx_media_hash ON media(hash);
	CREATE INDEX IF NOT EXISTS idx_media_path ON media(path);
	CREATE INDEX IF NOT EXISTS idx_media_status ON media(status);
	CREATE INDEX IF NOT EXISTS idx_media_last_viewed_at ON media(last_viewed_at);

	-- ==================== 媒体内容表 ====================
	CREATE TABLE IF NOT EXISTS media_content (
		media_id INTEGER NOT NULL,
		content_type TEXT NOT NULL,
		content TEXT NOT NULL DEFAULT '',
		updated_at TEXT NOT NULL DEFAULT (strftime('%Y-%m-%dT%H:%M:%fZ','now')),
		PRIMARY KEY (media_id, content_type),
		FOREIGN KEY (media_id) REFERENCES media(id) ON DELETE CASCADE
	);
	CREATE INDEX IF NOT EXISTS idx_media_content_type ON media_content(content_type);

	-- ==================== 标签字典 ====================
	CREATE TABLE IF NOT EXISTS tag (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL DEFAULT '',
		name_norm TEXT NOT NULL DEFAULT '',
		source TEXT NOT NULL DEFAULT 'manual',
		created_at TEXT NOT NULL DEFAULT (strftime('%Y-%m-%dT%H:%M:%fZ','now'))
	);
	CREATE UNIQUE INDEX IF NOT EXISTS uniq_tag_name_norm ON tag(name_norm);
	CREATE INDEX IF NOT EXISTS idx_tag_name ON tag(name);

	-- ==================== 媒体-标签关系 ====================
	CREATE TABLE IF NOT EXISTS media_tag (
		media_id INTEGER NOT NULL,
		tag_id INTEGER NOT NULL,
		PRIMARY KEY (media_id, tag_id),
		FOREIGN KEY (media_id) REFERENCES media(id) ON DELETE CASCADE,
		FOREIGN KEY (tag_id) REFERENCES tag(id) ON DELETE CASCADE
	);
	CREATE INDEX IF NOT EXISTS idx_media_tag_tag ON media_tag(tag_id);

	-- ==================== 分类字典 ====================
	CREATE TABLE IF NOT EXISTS category (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL DEFAULT '',
		parent_id INTEGER,
		level INTEGER NOT NULL DEFAULT 1,
		path TEXT NOT NULL DEFAULT '',
		sort INTEGER NOT NULL DEFAULT 0,
		created_at TEXT NOT NULL DEFAULT (strftime('%Y-%m-%dT%H:%M:%fZ','now')),
		FOREIGN KEY (parent_id) REFERENCES category(id) ON DELETE RESTRICT
	);
	CREATE INDEX IF NOT EXISTS idx_category_parent ON category(parent_id);
	CREATE INDEX IF NOT EXISTS idx_category_path ON category(path);

	-- ==================== 媒体-分类关系 ====================
	CREATE TABLE IF NOT EXISTS media_category (
		media_id INTEGER NOT NULL,
		category_id INTEGER NOT NULL,
		is_primary INTEGER NOT NULL DEFAULT 0,
		PRIMARY KEY (media_id, category_id),
		FOREIGN KEY (media_id) REFERENCES media(id) ON DELETE CASCADE,
		FOREIGN KEY (category_id) REFERENCES category(id) ON DELETE CASCADE
	);
	CREATE INDEX IF NOT EXISTS idx_media_category_category ON media_category(category_id);

	-- ==================== 用户行为流水 ====================
	CREATE TABLE IF NOT EXISTS media_behavior (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		media_id INTEGER NOT NULL,
		behavior_type TEXT NOT NULL DEFAULT '',
		behavior_value TEXT NOT NULL DEFAULT '{}',
		behavior_source TEXT NOT NULL DEFAULT 'manual',
		created_at TEXT NOT NULL DEFAULT (strftime('%Y-%m-%dT%H:%M:%fZ','now')),
		FOREIGN KEY (media_id) REFERENCES media(id) ON DELETE CASCADE
	);
	CREATE INDEX IF NOT EXISTS idx_behavior_media ON media_behavior(media_id);
	CREATE INDEX IF NOT EXISTS idx_behavior_type ON media_behavior(behavior_type);
	CREATE INDEX IF NOT EXISTS idx_behavior_created_at ON media_behavior(created_at);
	CREATE INDEX IF NOT EXISTS idx_behavior_type_created ON media_behavior(behavior_type, created_at);

	-- ==================== 搜索历史 ====================
	CREATE TABLE IF NOT EXISTS search_history (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		keyword TEXT NOT NULL DEFAULT '',
		keyword_norm TEXT NOT NULL DEFAULT '',
		use_count INTEGER NOT NULL DEFAULT 1,
		last_used_at TEXT NOT NULL DEFAULT (strftime('%Y-%m-%dT%H:%M:%fZ','now')),
		created_at TEXT NOT NULL DEFAULT (strftime('%Y-%m-%dT%H:%M:%fZ','now'))
	);
	CREATE UNIQUE INDEX IF NOT EXISTS uniq_search_keyword_norm ON search_history(keyword_norm);
	CREATE INDEX IF NOT EXISTS idx_search_last_used ON search_history(last_used_at);

	-- ==================== 扫描索引 ====================
	CREATE TABLE IF NOT EXISTS scanner_index (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		media_id INTEGER,
		file_path TEXT NOT NULL DEFAULT '',
		file_size INTEGER NOT NULL DEFAULT 0,
		modified_time TEXT NOT NULL DEFAULT (strftime('%Y-%m-%dT%H:%M:%fZ','now')),
		file_hash TEXT NOT NULL DEFAULT '',
		last_scan_at TEXT NOT NULL DEFAULT (strftime('%Y-%m-%dT%H:%M:%fZ','now')),
		FOREIGN KEY (media_id) REFERENCES media(id) ON DELETE SET NULL
	);
	CREATE UNIQUE INDEX IF NOT EXISTS uniq_scanner_path ON scanner_index(file_path);
	CREATE INDEX IF NOT EXISTS idx_scanner_hash ON scanner_index(file_hash);
	CREATE INDEX IF NOT EXISTS idx_scanner_modified ON scanner_index(modified_time);
	CREATE INDEX IF NOT EXISTS idx_scanner_media ON scanner_index(media_id);

	-- ==================== 全文搜索（FTS5） ====================
	CREATE VIRTUAL TABLE IF NOT EXISTS media_fts USING fts5(
		title, description, content='', tokenize='unicode61 remove_diacritic 2'
	);

	-- ==================== 触发器 ====================
	CREATE TRIGGER IF NOT EXISTS trg_media_updated_at AFTER UPDATE ON media
	BEGIN
		UPDATE media SET updated_at = strftime('%Y-%m-%dT%H:%M:%fZ','now') WHERE id = NEW.id;
	END;

	CREATE TRIGGER IF NOT EXISTS trg_media_fts_insert AFTER INSERT ON media
	BEGIN
		INSERT INTO media_fts(rowid, title, description) VALUES (NEW.id, NEW.title, NEW.description);
	END;

	CREATE TRIGGER IF NOT EXISTS trg_media_fts_delete AFTER DELETE ON media
	BEGIN
		INSERT INTO media_fts(media_fts, rowid, title, description) VALUES('delete', OLD.id, OLD.title, OLD.description);
	END;

	CREATE TRIGGER IF NOT EXISTS trg_media_fts_update AFTER UPDATE ON media
	BEGIN
		INSERT INTO media_fts(media_fts, rowid, title, description) VALUES('delete', OLD.id, OLD.title, OLD.description);
		INSERT INTO media_fts(rowid, title, description) VALUES (NEW.id, NEW.title, NEW.description);
	END;
	`

	_, err := db.Exec(schema)
	if err != nil {
		return fmt.Errorf("execute schema: %w", err)
	}

	return nil
}
