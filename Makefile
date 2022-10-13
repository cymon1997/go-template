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