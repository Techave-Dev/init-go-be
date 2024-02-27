ifneq ("$(wildcard .env)","")
	include .env
endif

mg-create:
	@migrate create -ext sql -dir database/migrations -seq $(name)

mg-up:
	@migrate -path database/migrations -database $(POSTGRES_URL) --verbose up
	@go run cmd/seed/main.go

mg-down:
	@migrate -path database/migrations -database $(POSTGRES_URL) --verbose down

seed:
	@go run cmd/seed/main.go

migrate:
	@go run cmd/migrate/main.go
	@go run cmd/seed/main.go

sqlc:
	@rm -rf internal/repo/psql/
	@sqlc generate

build:
	@migrate -path database/migrations -database $(POSTGRES_URL) --verbose up
	@rm -rf internal/repo/psql/
	@sqlc generate
	@go build -o api cmd/api/main.go

start:
	@migrate -path database/migrations -database $(POSTGRES_URL) --verbose up
	@rm -rf internal/repo/psql/
	@sqlc generate
	@go build -o api cmd/api/main.go
	@./api

reset:
	@migrate -path database/migrations -database $(POSTGRES_URL) --verbose down
	@migrate -path database/migrations -database $(POSTGRES_URL) --verbose up
	@go run cmd/seed/main.go
	@rm -rf internal/repo/psql/
	@sqlc generate

dev:
	@rm -rf internal/repo/psql/
	@sqlc generate
	@air .
