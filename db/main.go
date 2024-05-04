package main

import (
	"context"
	"flag"
	"log"
	"os"
	"strings"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

var (
	flags = flag.NewFlagSet("goose", flag.ExitOnError)
	dir   = flags.String("dir", ".", "directory with migration files")
	db    = flags.String("db", "", "database connection string")
)

func main() {
	flags.Parse(os.Args[1:])
	args := flags.Args()

	if len(args) < 2 {
		flags.Usage()
		return
	}

	migrationDir, dbstring, command := args[0], args[1], args[2]

	parsedMigrationDir := strings.Split(migrationDir, "=")[1]
	parsedDbString := strings.SplitN(dbstring, "=", 2)[1]

	db, err := goose.OpenDBWithDriver("postgres", parsedDbString)
	if err != nil {
		log.Fatalf("goose: failed to open DB: %v\n", err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("goose: failed to close DB: %v\n", err)
		}
	}()

	arguments := []string{}
	if len(args) > 3 {
		arguments = append(arguments, args[3:]...)
	}

	ctx := context.Background()
	if err := goose.RunContext(ctx, command, db, parsedMigrationDir, arguments...); err != nil {
		log.Fatalf("goose %v: %v", command, err)
	}
}
