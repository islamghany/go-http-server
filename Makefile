# include dev.env

# Check to see if we can use ash, in Alpine images, or default to BASH.
SHELL_PATH = /bin/ash
SHELL = $(if $(wildcard $(SHELL_PATH)),/bin/ash,/bin/bash)

WEB_APP_VERSION = 0.0.1
WEB_APP_NAME = app_name

MIGRATION_PATH = ./internal/db/migrations
DATABASE_URL = postgres://root:secret@localhost:5321/blogdb?sslmode=disable


run:
	go run cmd/app_name/main.go

run-migrate:
	go run cmd/admin/main.go migrate

run-migrate-down:
	go run cmd/admin/main.go migrate-down
run-init:
	go run cmd/admin/main.go migrate,seed

build-dev:
	go build  -ldflags "-X main.version=$(WEB_APP_VERSION) -X main.build=development" -o bin/app_name cmd/app_name/main.go

build-prod:
	go build -ldflags "-X main.version=$(WEB_APP_VERSION) -X main.build=production" -o bin/app_name cmd/app_name/main.go 





migrate/new:
	@migrate create -ext sql -dir $(MIGRATION_PATH) -seq ${name}

# Migrate to the latest version of the database.
migrate/up:
	@migrate -path $(MIGRATION_PATH) -database $(DATABASE_URL)  -verbose up

# Migrate to the previous version of the database.
migrate/down:
	@migrate -path $(MIGRATION_PATH) -database $(DATABASE_URL) -verbose down

# Migrate to the latest version of the database.
migrate/redo:
	@migrate -path $(MIGRATION_PATH) -database $(DATABASE_URL) -verbose redo
