#!/bin/bash

DBSTRING="host=$POSTGRES_HOST user=$POSTGRES_USERNAME password=$POSTGRES_PASSWORD dbname=$POSTGRES_DB_NAME sslmode=$POSTGRES_SSL_MODE"

goose postgres "$DBSTRING" up