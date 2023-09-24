#!/usr/bin/env bash

# Read database schema and auto-generate models
# Copy these code block for each separated database
EXPORT DB_DEFAULT_NAME=go_template_test
EXPORT DB_DEFAULT_CODE_PATH=./core/outbound/db
EXPORT DB_DEFAULT_DSN="postgres://go_template:root@127.0.0.1:15432/$DB_DEFAULT_NAME?sslmode=disable"
gen --sqltype=postgres --connstr "$DSN" --database $DB_DEFAULT_NAME --db --gorm --guregu --overwrite --out $DB_DEFAULT_CODE_PATH
goimports -w "$DB_DEFAULT_CODE_PATH/model"