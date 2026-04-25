APP_NAME := daily-news
BIN_DIR := bin

.PHONY: dev run-backend run-frontend build build-web build-backend clean \
	build-linux build-linux-arm64 build-macos build-macos-arm64 build-windows build-all

dev:
	@echo "Starting backend (:8080) and frontend (:5173)..."
	@bash -c 'set -m; go run ./cmd/app & npm run dev; kill %1 2>/dev/null || true'

run-backend:
	go run ./cmd/app

run-frontend:
	npm run dev

build: build-web build-backend

build-web:
	npm run build

build-backend:
	mkdir -p $(BIN_DIR)
	CGO_ENABLED=0 go build -o $(BIN_DIR)/$(APP_NAME) ./cmd/app

build-linux: build-web
	mkdir -p $(BIN_DIR)
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o $(BIN_DIR)/$(APP_NAME)-linux-amd64 ./cmd/app

build-linux-arm64: build-web
	mkdir -p $(BIN_DIR)
	GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -o $(BIN_DIR)/$(APP_NAME)-linux-arm64 ./cmd/app

build-macos: build-web
	mkdir -p $(BIN_DIR)
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -o $(BIN_DIR)/$(APP_NAME)-darwin-amd64 ./cmd/app

build-macos-arm64: build-web
	mkdir -p $(BIN_DIR)
	GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 go build -o $(BIN_DIR)/$(APP_NAME)-darwin-arm64 ./cmd/app

build-windows: build-web
	mkdir -p $(BIN_DIR)
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -o $(BIN_DIR)/$(APP_NAME)-windows-amd64.exe ./cmd/app

build-all: build-linux build-linux-arm64 build-macos build-macos-arm64 build-windows

clean:
	rm -rf $(BIN_DIR) site/dist
