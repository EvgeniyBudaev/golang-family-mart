#!/usr/bin/env bash
set -x
set +e
curl -L https://github.com/golang-migrate/migrate/releases/download/v4.16.2/migrate.linux-amd64.tar.gz | tar xvz
mv migrate scripts/migrate
chmod +x scripts/migrate
rm README.MD