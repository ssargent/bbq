package migration

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/pgx/v5"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/ssargent/bbq/internal/config"
)

type DB struct {
	migrations string
}

func NewDB(path string) *DB {
	return &DB{
		migrations: path,
	}
}

func (d *DB) Run(cfg *config.Config) error {
	db, err := sql.Open("pgx", cfg.Database.Uri())
	if err != nil {
		return fmt.Errorf("sql.Open: %w", err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("postres.WithInstance: %w", err)
	}

	migrations, err := migrate.NewWithDatabaseInstance(d.migrations, "postgres", driver)
	if err != nil {
		return fmt.Errorf("migrate.NewWithDatabaseInstance: %w", err)
	}

	if err := migrations.Up(); err != nil {
		return fmt.Errorf("migrations.Up: %w", err)
	}

	return nil
}
