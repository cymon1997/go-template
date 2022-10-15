#!/usr/bin/env bash

DSN='postgres://go_template:root@127.0.0.1:15432/go_template_test?sslmode=disable'

gen --sqltype=postgres --connstr "$DSN" \
 --database go_template_test --db --gorm --guregu --overwrite --out ./core/outbound/db

goimports -w ./core/outbound/db/model