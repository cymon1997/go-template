#!/usr/bin/env bash
export NOW=$(shell date +"%F %T")

## Development
update-dep:
	@printf "%s %s > UPDATE DEPENDENCIES ... \n" $(NOW)
	@go mod tidy
	@echo 'done'

## Testing Utility
lint:
	@golangci-lint run --modules-download-mode=readonly

utest:
	@go test ./...

test: lint utest

## Running app
init-infra:
	@docker-compose up -d

build:
	@printf "%s %s > BUILD APP BINARY ... \n" $(NOW)
	@go build -race -o app ./cmd/app/main.go
	@echo 'done'

run:
	@printf "%s %s > RUN APP ... \n" $(NOW)
	@ENV=LOCAL MallocNanoZone=0 ./app

quick: build run

## Development Tools
install-tools:
	@printf "%s %s > INSTALL TOOLS ... \n" $(NOW)
	@./scripts/install-tools.sh
	@echo 'done'

connect-db:
	@psql -h 127.0.0.1 -p 15432 -U go_template go_template_test

# ex: create-migration name=create_table_sample
create-migration:
	@migrate create -dir ./files/db/migrations -ext sql $(name)

# ex: create-seed name=insert_sample
create-seed:
	@migrate create -dir ./files/db/seeds -ext sql $(name)

generate-schema:
	@./scripts/generate-schema.sh
