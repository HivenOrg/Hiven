#!/usr/bin/env bash

# Don't run manually, docker compose will run it

PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -U $DB_USER -d postgres -c "DROP DATABASE IF EXISTS $TEST_DB_NAME;"
PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -U $DB_USER -d postgres -c "CREATE DATABASE $TEST_DB_NAME;"
