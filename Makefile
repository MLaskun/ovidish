## help: you are here
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

## tidy: format all .go files and tidy module dependencies
.PHONY: tidy
tidy:
	@echo 'Formatting .go files...'
	go fmt ./...
	@echo 'Tidying module dependencies'
	go mod tidy

## db/migrations/new name=$1: create new database migration
.PHONY: db/migrations/new
db/migrations/new:
	@echo 'Creating migration files for product with name: ${name}...'
	migrate create -seq -ext=.sql -dir=./migrations ${name}

## db/migratuins/up: apply all up database migrations
.PHONY: db/migrations/up
db/migrations/up:
	@echo 'Running up migrations...'
	migrate -path=./migrations -database=postgres://postgres:postgres@127.0.0.1:5432/ovidish?sslmode=disable up
