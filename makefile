# Makefile

# Docker
.PHONY: build up down logs

build:
	docker-compose build

up:
	docker-compose up -d

down:
	docker-compose down

logs:
	docker-compose logs -f

migrate-up:
	docker-compose exec app go run migrations/up.go

migrate-down:
	docker-compose exec app go run migrations/down.go
