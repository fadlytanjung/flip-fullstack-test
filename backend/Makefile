# ---- build, run, and test

.PHONY: run dev build lint format test tidy install help coverage

help:
	@echo "Available commands:"
	@echo "  make run          - Run the server (standard)"
	@echo "  make dev          - Run the server with hot reload (air)"
	@echo "  make build        - Build the application binary"
	@echo "  make lint         - Run linter"
	@echo "  make format       - Format code"
	@echo "  make test         - Run tests"
	@echo "  make coverage     - Run tests with coverage report"
	@echo "  make tidy         - Tidy dependencies"
	@echo "  make install      - Install development tools"

run:
	go run ./cmd/server

dev:
	$(HOME)/go/bin/air

build:
	go build -o ./app ./cmd/server

# ---- code quality

lint:
	golangci-lint run ./...

format:
	go fmt ./...

test:
	go test ./...

coverage:
	@mkdir -p coverage
	go test -coverpkg=./... -covermode=atomic -coverprofile=coverage/coverage.out.tmp ./...
	cat coverage/coverage.out.tmp | grep -v -E "cmd|mocks" > coverage/coverage.out
	go tool cover -html=coverage/coverage.out -o coverage/coverage.html
	@echo ""
	@echo "==================COVERAGE SUMMARY=================="
	@echo ""
	@go tool cover -func coverage/coverage.out | grep total | awk '{print substr($$3, 1, length($$3)-1)}' | xargs printf "Total Coverage: %s%%\n"
	@echo ""
	@echo "HTML Report: coverage/coverage.html"
	@echo ""

# ---- dependencies

tidy:
	go mod tidy

install:
	go mod download
	go install github.com/cosmtrek/air@v1.51.0
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.57.2