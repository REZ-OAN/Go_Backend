#!/bin/sh

set -e 

echo "executing database migration"
/app/migrate -path /app/migrations -database "$DB_SOURCE" -verbose up

echo "starting the app"
exec "$@"