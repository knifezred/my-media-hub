package service

import (
	"crypto/sha256"
	"database/sql"
	"fmt"
	"io"
	"io/fs"
	"my-media-hub/backend/internal/model"
	"my-media-hub/backend/internal/repository"
	"my-media-hub/backend/internal/search"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type scanTask struct {
	mu        sync.Mutex
	running   bool
	processed int
	total     int
}

var currentScan = &scanTask{}

type ScannerService struct {
	scannerRepo *repository.ScannerRepository
	mediaRepo   *repository.MediaRepository
	index       *search.Index
}

func NewScannerService(db *sql.DB, index *search.Index) *ScannerService {
	return &ScannerService{
		scannerRepo: repository.NewScannerRepository(db),
		mediaRepo:   repository.NewMediaRepository(db),
		index:       index,
	}
}

func (s *ScannerService) Start(dirs []string) error {
	if len(dirs) == 0 {
		return fmt.Errorf("no scan directories configured")
	}

	currentScan.mu.Lock()
	if currentScan.running {
		currentScan.mu.Unlock()
		return fmt.Errorf("scan already in progress")
	}
	currentScan.running = true
	currentScan.processed = 0
	currentScan.total = 0
	currentScan.mu.Unlock()

	go s.run(dirs)

	return nil
}

func (s *ScannerService) Status() model.ScannerStatus {
	currentScan.mu.Lock()
	defer currentScan.mu.Unlock()

	progress := 0.0
	if currentScan.total > 0 {
		progress = float64(currentScan.processed) / float64(currentScan.total) * 100
	}

	return model.ScannerStatus{
		Running:   currentScan.running,
		Processed: currentScan.processed,
		Total:     currentScan.total,
		Progress:  progress,
	}
}

type fileEntry struct {
	path string
	size int64
}

func (s *ScannerService) run(dirs []string) {
	defer func() {
		currentScan.mu.Lock()
		currentScan.running = false
		currentScan.mu.Unlock()
	}()

	var allFiles []fileEntry
	for _, dir := range dirs {
		filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return nil
			}
			if d.IsDir() {
				return nil
			}
			info, err := d.Info()
			if err != nil {
				return nil
			}
			allFiles = append(allFiles, fileEntry{path: path, size: info.Size()})
			return nil
		})
	}

	currentScan.mu.Lock()
	currentScan.total = len(allFiles)
	currentScan.mu.Unlock()

	for _, f := range allFiles {
		if err := s.processFile(f); err != nil {
		}

		currentScan.mu.Lock()
		currentScan.processed++
		currentScan.mu.Unlock()
	}
}

func (s *ScannerService) processFile(f fileEntry) error {
	ext := strings.ToLower(filepath.Ext(f.path))
	mediaType := detectMediaType(ext)
	if mediaType == "" {
		return nil
	}

	title := extractTitle(f.path)

	hash, err := computeHash(f.path)
	if err != nil {
		return fmt.Errorf("compute hash: %w", err)
	}

	exists, err := s.scannerRepo.ExistsByHash(hash)
	if err != nil {
		return fmt.Errorf("check hash: %w", err)
	}
	if exists {
		return nil
	}

	coverPath := ""
	if mediaType == "image" {
		coverPath = f.path
	}

	id, err := s.mediaRepo.Insert(&model.Media{
		MediaType: mediaType,
		Title:     title,
		Path:      f.path,
		Hash:      hash,
		Size:      f.size,
		CoverPath: coverPath,
	})
	if err != nil {
		return fmt.Errorf("insert media: %w", err)
	}

	if err := s.index.IndexMedia(uint64(id), title, ""); err != nil {
		return fmt.Errorf("index media: %w", err)
	}

	return nil
}

func detectMediaType(ext string) string {
	switch ext {
	case ".jpg", ".jpeg", ".png", ".gif", ".webp", ".bmp", ".svg", ".ico":
		return "image"
	case ".mp4", ".mkv", ".avi", ".mov", ".wmv", ".flv", ".webm", ".ts", ".mts":
		return "video"
	case ".txt", ".md", ".epub", ".mobi", ".azw3", ".fb2":
		return "novel"
	default:
		return ""
	}
}

func extractTitle(path string) string {
	base := filepath.Base(path)
	ext := filepath.Ext(base)
	return strings.TrimSuffix(base, ext)
}

func computeHash(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
