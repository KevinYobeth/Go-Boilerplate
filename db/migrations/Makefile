DB_DRIVER=postgres
DB_STRING="host=localhost port=5432 user=postgres password=postgres dbname=library sslmode=disable"

goose_up:
	env GOOSE_DRIVER=$(DB_DRIVER) GOOSE_DBSTRING=$(DB_STRING) goose up

goose_down:
	env GOOSE_DRIVER=$(DB_DRIVER) GOOSE_DBSTRING=$(DB_STRING) goose down

goose_status:
	env GOOSE_DRIVER=$(DB_DRIVER) GOOSE_DBSTRING=$(DB_STRING) goose status