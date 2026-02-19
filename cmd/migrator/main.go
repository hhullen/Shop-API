package main

import (
	"context"
	"database/sql"
	"log"
	"os"

	"shopapi/internal/clients/postgres"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

const (
	migrationsDirPostgres = "./migrations/postgres"
	cmdUp                 = "up"
	cmdDown               = "down"
	cmdStatus             = "status"
)

var executors = map[string]func(db *sql.DB, dir string, opts ...goose.OptionsFunc) error{
	cmdUp:     goose.Up,
	cmdDown:   goose.Down,
	cmdStatus: goose.Status,
}

func main() {
	defer func() {
		if p := recover(); p != nil {
			log.Fatalf("%v", p)
		}
	}()

	if len(os.Args) < 2 {
		log.Fatalf("migration action is required: %s/%s/%s", cmdUp, cmdDown, cmdStatus)
	}
	command := os.Args[1]

	goose.SetBaseFS(nil)

	MigratePostgres(command)
}

func MigratePostgres(command string) {
	ctx := context.Background()
	db, err := postgres.NewSQLConn(ctx)
	if err != nil {
		log.Fatalf("failed to open DB: %v", err)
	}

	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatalf("failed to set dialect: %v", err)
	}

	ExecMigration(db, command, migrationsDirPostgres)

	ctx.Done()
}

func ExecMigration(db *sql.DB, command, migrationsDir string) {
	executor, exists := executors[command]
	if !exists {
		log.Fatalf("Wrong comand send: %s. Required: %s/%s/%s", command, cmdUp, cmdDown, cmdStatus)
	}

	if err := executor(db, migrationsDir); err != nil {
		log.Fatalf("failed to apply migrations: %v", err)
	}

	log.Printf("%s successfully migrated\n", migrationsDir)
}
