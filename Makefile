.PHONY: run-backend run-frontend build-backend build-frontend build-all test lint lint-backend lint-frontend clean swag pack help

# ========== Paths ==========
BACKEND_DIR  := backend
FRONTEND_DIR := web

# ========== Build Flags ==========
BUILD_NUM ?= 1

# ========== Windows Commands ==========
RM := powershell -Command "Remove-Item -Recurse -Force -ErrorAction SilentlyContinue"

# ========== Development ==========
run-backend:
	@echo "[INFO] Starting backend server..."
	@echo "[INFO] API: http://localhost:8080"
	cd $(BACKEND_DIR) && go run ./cmd/

run-frontend:
	@echo "[INFO] Starting frontend dev server..."
	@echo "[INFO] App:  http://localhost:5173"
	cd $(FRONTEND_DIR) && pnpm run dev

# ========== Build ==========
build-backend:
	@echo "[INFO] Building Go backend..."
	cd $(BACKEND_DIR) && go build -o bin/media-hub ./cmd/
	@echo "[OK] Backend built: $(BACKEND_DIR)/bin/media-hub"

build-frontend:
	@echo "[INFO] Building Vue frontend..."
	cd $(FRONTEND_DIR) && pnpm install && pnpm run build
	@echo "[OK] Frontend built: $(FRONTEND_DIR)/dist/"

build-all: build-backend build-frontend
	@echo "[OK] All builds complete"

# ========== Docs ==========
swag:
	@echo "[INFO] Generating Swagger docs..."
	cd $(BACKEND_DIR) && swag init -g ./cmd/main.go -o ./docs --parseDependency --parseInternal
	@echo "[OK] Swagger docs generated: $(BACKEND_DIR)/docs/"

# ========== Pack ==========
pack: build-all
	@echo "[INFO] Running ugcli pack (build $(BUILD_NUM))..."
	ugcli pack --arch amd64 --build $(BUILD_NUM)
	@echo "[OK] Pack completed"

# ========== Quality ==========
test:
	cd $(BACKEND_DIR) && go test -v -count=1 ./...

lint-backend:
	cd $(BACKEND_DIR) && go vet ./...

lint-frontend:
	cd $(FRONTEND_DIR) && pnpm run typecheck

lint: lint-backend lint-frontend

# ========== Clean ==========
clean:
	@echo "[INFO] Cleaning build artifacts..."
	$(RM) $(BACKEND_DIR)/bin
	$(RM) $(FRONTEND_DIR)/dist
	@echo "[OK] Cleaned"

# ========== Help ==========
help:
	@echo "Available commands:"
	@echo ""
	@echo "Development:"
	@echo "  make run-backend   - Start backend (http://localhost:8080)"
	@echo "  make run-frontend  - Start frontend (http://localhost:5173)"
	@echo ""
	@echo "Build:"
	@echo "  make build-backend   - Build backend binary"
	@echo "  make build-frontend  - Build frontend assets"
	@echo "  make build-all       - Build both backend and frontend"
	@echo ""
	@echo "Docs:"
	@echo "  make swag            - Generate Swagger docs"
	@echo ""
	@echo "Pack:"
	@echo "  make pack BUILD_NUM=1 - Pack with build version"
	@echo ""
	@echo "Quality:"
	@echo "  make test         - Run Go tests"
	@echo "  make lint         - Run linters (backend + frontend)"
	@echo "  make lint-backend - Run go vet"
	@echo "  make lint-frontend- Run vue-tsc typecheck"
	@echo ""
	@echo "Other:"
	@echo "  make clean - Remove build artifacts"
