#!/bin/sh
CURRENT_DIR="$(pwd)"

export DATABASE_URL="postgres://postgres:hwhwhwlol@localhost:5432/tanyaustadz_db?sslmode=disable" \
export MIGRATION_PATH="${CURRENT_DIR}/src/tanyaustadz/migration/pgsql" \
# go run migration/main.go migration:status
# go run migration/main.go migration:down
go run migration/main.go migration:up
# go run migration/main.go migration:create create_table_orders --table=orders