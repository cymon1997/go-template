#!/usr/bin/env bash

set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    CREATE DATABASE go_template_test;
    GRANT ALL PRIVILEGES ON DATABASE go_template_test TO go_template;
EOSQL