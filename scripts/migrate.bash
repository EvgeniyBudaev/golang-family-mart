#!/usr/bin/env bash
set -x
set +e
source .env
MIGRATE="./scripts/migrate"
#-source ./migrations
$MIGRATE  -path ./migrations -database $DATABASE_URL up