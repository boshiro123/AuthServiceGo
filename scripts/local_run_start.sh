#!/bin/bash

# docker network create legal-docs-ai-net || true

docker compose \
  --env-file compose/env/local/.env \
  -f compose/base/docker-compose.db.yml \
  up -d

./scripts/wait-for-it.sh localhost:5432 --timeout=60 --strict -- echo "Postgres is up"

docker compose \
  --env-file compose/env/local/.env \
  -f compose/base/docker-compose.db.yml \
  -f compose/env/local/docker-compose.jobs.yml up --build -d

services=(
  "legal-docs-ai-migrations"
)

for service in "${services[@]}"; do
  container_id=$(docker ps -a -q -f name="$service")

  if [ -n "$container_id" ]; then
    docker wait "$container_id"
  else
    echo "Container for the $service service has already been stopped."
  fi

  exit_code=$(docker inspect "$container_id" --format='{{.State.ExitCode}}')
  if [ "$exit_code" -ne 0 ]; then
    echo "The $service service has exited with an error. Exit code: $exit_code"
    exit 1
  else
    echo "The $service service has successfully exited. Exit code: $exit_code"
  fi
done

docker compose \
  --env-file compose/env/local/.env \
  -f compose/base/docker-compose.db.yml \
  -f compose/env/local/docker-compose.override.yml up --build -d