/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	_ "github.com/golang-migrate/migrate/v4/database/pgx/v5"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/spf13/cobra"
	"github.com/ssargent/bbq/internal/config"
	"github.com/ssargent/bbq/internal/migration"
)

type migrateApp struct {
	migrations string
}

func init() { //nolint:gochecknoinits // required by cobra
	m := &migrateApp{}
	cmd := &cobra.Command{
		Use:   "migrate",
		Short: "Migrate the database",
		Long:  "Migrate the database",
		Run: func(cmd *cobra.Command, args []string) {
			err := m.run()
			if err != nil {
				fmt.Println(err)
			}
		},
	}

	cmd.Flags().StringVarP(&m.migrations, "migrations", "m", "file://../migrations", "path to the migrations")
	rootCmd.AddCommand(cmd)
}

func (m *migrateApp) run() error {
	cfg := config.Config{}

	if err := godotenv.Load(runEnvFile); err != nil {
		return fmt.Errorf("godotenv.Load: %w", err)
	}

	if err := envconfig.Process("bbqd", &cfg); err != nil {
		return err
	}

	migrate := migration.NewDB(m.migrations)
	if err := migrate.Run(&cfg); err != nil {
		return fmt.Errorf("migrate.Run: %w", err)
	}
	return nil
}
