.PHONY: help install build run dev test clean fmt lint docker-up docker-down migrate-up migrate-down

help:
	@echo "Available commands:"
	@echo "  make install      - Install dependencies"
	@echo "  make build        - Build the application"
	@echo "  make run          - Run the application"
	@echo "  make dev          - Run in development mode"
	@echo "  make test         - Run tests"
	@echo "  make clean        - Clean build artifacts"
	@echo "  make fmt          - Format code"
	@echo "  make lint         - Run linter"
	@echo "  make docker-up    - Start local Postgres"
	@echo "  make docker-down  - Stop local Postgres"
	@echo "  make migrate-up   - Run pending migrations"
	@echo "  make migrate-down - Rollback last migration"

install:
	go mod download
	go mod tidy

build:
	go build -o bin/api cmd/api/main.go

run: build
	./bin/api

dev:
	go run cmd/api/main.go

test:
	go test -v ./...

clean:
	rm -rf bin/

fmt:
	go fmt ./...

lint:
	golangci-lint run

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

migrate-up:
	go run cmd/migrate/main.go up

migrate-down:
	go run cmd/migrate/main.go down
