#!/usr/bin/env bash
export NOW=$(shell date +"%F %T")

## Pre-run
update-dep:
	@printf "%s %s >> UPDATE DEPENDENCIES ... \n" $(NOW)
	@go mod tidy
	@echo 'done'

## Test
lint:
	@golangci-lint run --modules-download-mode=readonly

utest:
	@go test ./...

test: lint utest

# ex: create-migration name=create_table_sample
create-migration:
	@migrate create -dir ./files/db/migrations -ext sql $(name)

# ex: create-seed name=insert_sample
create-seed:
	@migrate create -dir ./files/db/seeds -ext sql $(name)

generate-schema:
	@./scripts/generate-schema.sh
