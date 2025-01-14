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
