ENV_FILE := .env
APP_DIR := ./...
CMD_DIR := ./cmd/api
BIN_DIR := bin

ifneq ("$(wildcard $(ENV_FILE))","")
include $(ENV_FILE)
export
endif

.PHONY: help init add add-tool run test cov lint format check clean reset

help:
	@echo "Available commands:"
	@echo "  make init                        - Download dependencies and prepare folders"
	@echo "  make add p=github.com/pkg/errors - Add a Go module dependency"
	@echo "  make add-tool p=...              - Install a Go-based development tool"
	@echo "  make run                         - Run the main program"
	@echo "  make test                        - Run tests"
	@echo "  make cov                         - Run tests with coverage"
	@echo "  make lint                        - Run go vet"
	@echo "  make format                      - Format code"
	@echo "  make check                       - Run format + lint + test"
	@echo "  make clean                       - Remove temporary files"
	@echo "  make reset                       - Recreate local artifacts"

init:
	@mkdir -p $(BIN_DIR) logs
	go mod tidy

add:
	@test -n "$(p)" || (echo "Use: make add p=module_path" && exit 1)
	go get $(p)

run:
	go run $(CMD_DIR)

test:
	APP_ENV=test go test -v $(APP_DIR)

cov:
	APP_ENV=test go test -coverprofile=coverage.out $(APP_DIR)
	go tool cover -func=coverage.out

lint:
	go vet $(APP_DIR)

format:
	@files=$$(find . -name "*.go" -not -path "./vendor/*"); \
	if [ -n "$$files" ]; then gofmt -w $$files; fi

check: format lint test

clean:
	rm -rf $(BIN_DIR) logs coverage.out

reset: clean init
