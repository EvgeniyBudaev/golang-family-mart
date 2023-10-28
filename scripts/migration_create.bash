#!/usr/bin/env bash
set -x
set +e
source .env
MIGRATE="./scripts/migrate"
#$MIGRATE -source ./migrations -database $DATABASE_URL -ext=sql create migration
$MIGRATE create -ext sql -dir ./migrations/ -seq migration