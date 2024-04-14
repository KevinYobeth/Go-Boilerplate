include .env

DB_DRIVER=postgres
DB_STRING="host=${POSTGRES_HOST} port=${POSTGRES_PORT} user=${POSTGRES_USERNAME} password=${POSTGRES_PASSWORD} dbname=${POSTGRES_DB_NAME} sslmode=${POSTGRES_SSL_MODE}"
MIGRATION_DIR=db/migrations

db_up:
	GOOSE_DRIVER=$(DB_DRIVER) GOOSE_DBSTRING=$(DB_STRING) goose -dir=${MIGRATION_DIR} up

db_down:
	GOOSE_DRIVER=$(DB_DRIVER) GOOSE_DBSTRING=$(DB_STRING) goose -dir=${MIGRATION_DIR} down

db_status:
	GOOSE_DRIVER=$(DB_DRIVER) GOOSE_DBSTRING=$(DB_STRING) goose -dir=${MIGRATION_DIR} status

db_create:
	GOOSE_DRIVER=$(DB_DRIVER) GOOSE_DBSTRING=$(DB_STRING) goose -dir=${MIGRATION_DIR} create $(filter-out $@,$(MAKECMDGOALS)) sql

openapi:
	@./scripts/openapi-http.sh

proto:
	@./scripts/proto.sh	

run:
	SERVER_TYPE=$(filter-out $@,$(MAKECMDGOALS)) go run .