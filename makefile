include .env

DB_STRING="host=${POSTGRES_HOST} port=${POSTGRES_PORT} user=${POSTGRES_USERNAME} password=${POSTGRES_PASSWORD} dbname=${POSTGRES_DB_NAME} sslmode=${POSTGRES_SSL_MODE}"
MIGRATION_DIR=db/migrations

db_up:
	./migrate dir=${MIGRATION_DIR} db=${DB_STRING} up

db_down:
	./migrate dir=${MIGRATION_DIR} db=${DB_STRING} down

db_status:
	./migrate dir=${MIGRATION_DIR} db=${DB_STRING} status

db_create:
	./migrate dir=${MIGRATION_DIR} db=${DB_STRING} create $(filter-out $@,$(MAKECMDGOALS)) sql

.PHONY: openapi
openapi:
	@./scripts/openapi-http.sh

.PHONY: proto
proto:
	@./scripts/proto.sh	

.PHONY: run
run:
	SERVER_TYPE=$(filter-out $@,$(MAKECMDGOALS)) go run .