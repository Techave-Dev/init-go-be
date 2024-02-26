ifneq ("$(wildcard .env)","")
	include .env
endif

mg-create:
	@migrate create -ext sql -dir database/migrations -seq $(name)

mg-up:
	@migrate -path database/migrations -database $(POSTGRES_URL) --verbose up

mg-down:
	@migrate -path database/migrations -database $(POSTGRES_URL) --verbose down

seed:
	@go run cmd/seed/main.go

migrate:
	@go run cmd/migrate/main.go

sqlc:
	@rm -rf internal/repo
	@sqlc generate

build:
	@migrate -path database/migrations -database $(POSTGRES_URL) --verbose up
	@sqlc generate
	@go build -o api cmd/api/main.go

start:
	@migrate -path database/migrations -database $(POSTGRES_URL) --verbose up
	@sqlc generate
	@go build -o api cmd/api/main.go
	@./api

dev:
	@sqlc generate
	@air .
