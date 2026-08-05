#!/usr/bin/env bash

set -e

DB_URL="postgres://postgres:postgres@localhost:5432/chirpy?sslmode=disable"

goose -dir sql/schema postgres "$DB_URL" down
sleep 1
goose -dir sql/schema postgres "$DB_URL" up

sqlc generate

echo "Database reset and SQLC regenerated."