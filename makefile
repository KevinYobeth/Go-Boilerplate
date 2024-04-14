DB_DRIVER=postgres
DB_STRING="host=localhost port=5432 user=postgres password=postgres dbname=boilerplate sslmode=disable"
MIGRATION_DIR=db/migrations

db_up:
	env GOOSE_DRIVER=$(DB_DRIVER) GOOSE_DBSTRING=$(DB_STRING) goose -dir=${MIGRATION_DIR} up

db_down:
	env GOOSE_DRIVER=$(DB_DRIVER) GOOSE_DBSTRING=$(DB_STRING) goose -dir=${MIGRATION_DIR} down

db_status:
	env GOOSE_DRIVER=$(DB_DRIVER) GOOSE_DBSTRING=$(DB_STRING) goose -dir=${MIGRATION_DIR} status

db_create:
	env GOOSE_DRIVER=$(DB_DRIVER) GOOSE_DBSTRING=$(DB_STRING) goose -dir=${MIGRATION_DIR} create $(filter-out $@,$(MAKECMDGOALS)) sql