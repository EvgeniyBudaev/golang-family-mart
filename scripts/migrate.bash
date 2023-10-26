#!/usr/bin/env bash
set -x
set +e
source .env
MIGRATE="./scripts/migrate"
$MIGRATE -source ./migrations -database $DATABASE_URL up