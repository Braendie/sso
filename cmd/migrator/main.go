package main

import (
	"errors"
	"flag"
	"fmt"

	"github.com/Braendie/sso/internal/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	var migrationsTable string
	flag.StringVar(&migrationsTable, "migrations-table", "migrations", "name of migrations")

	cfg := config.MustLoad()

	var migrationsPath string
	if migrationsTable == "migrations_test" {
		migrationsPath = cfg.MigrationsTestPath
	} else {
		migrationsPath = cfg.MigrationsPath
	}

	m, err := migrate.New(
		"file://"+migrationsPath,
		fmt.Sprintf("sqlite3://%s?x-migrations-table=%s", cfg.StoragePath, migrationsTable),
	)
	if err != nil {
		panic(err)
	}

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("no migrations to update")

			return
		}

		panic(err)
	}

	fmt.Println("migrations applied successfully")
}
