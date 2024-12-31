#!/bin/bash
set -e

export PGPASSWORD="$POSTGRES_PASSWORD"

psql -v ON_ERROR_STOP=1 --host=test-database --port=5432 --username "$POSTGRES_USER" --dbname "postgres" <<-EOSQL
    DROP SCHEMA IF EXISTS bbb_voting_schema_test CASCADE;
    CREATE SCHEMA bbb_voting_schema_test;
EOSQL