/*
Copyright © 2023 Scott Sargent <scott.sargent@gmail.com>
*/
package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq" // <------------ here
	"github.com/patrickmn/go-cache"
	"github.com/spf13/cobra"
	"github.com/ssargent/bbq/cmd/bbq/internal"
	"github.com/ssargent/bbq/internal/config"
	"go.uber.org/zap"
)

// serveCmd represents the serve command.
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "BBQ Server Daemon",
	Long: `BBQ Server Daemon provides the full backend for the bbq project.  
	
	 - Quic Data Streaming Service
	 - Management API
	 - gRPC API`,
	Run: func(cmd *cobra.Command, args []string) {

		logger, err := zap.NewProduction()
		if err != nil {
			log.Fatalf("zap.NewProduction(): %s", err.Error())
		}

		api, err := server(logger)
		if err != nil {
			log.Fatalf("server: %w", err)
		}

		if err := api.ListenAndServe(); err != nil {
			log.Fatalf("ListenAndServe: %w", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func server(logger *zap.Logger) (*internal.API, error) {
	// in production we don't want to use something.env type files.
	// get our values from the orchestration layer itself.
	// but in development, its super useful.
	// so check to see if file even exists before trying to do the godotenv
	if _, err := os.Stat(runEnvFile); err == nil {
		if err := godotenv.Load(runEnvFile); err != nil {
			return nil, fmt.Errorf("godotenv.Load: %w", err)
		}
	}

	var cfg config.Config
	if err := envconfig.Process("bbqd", &cfg); err != nil {
		return nil, fmt.Errorf("envconfig.Process: %w", err)
	}

	explainConfig(&cfg)

	db, safeDb, err := database(&cfg)
	if err != nil {
		return nil, fmt.Errorf("database: %w", err)
	}

	fmt.Printf("Connecting to %s\n", safeDb)

	cache := cache.New(cfg.Cache.DefaultExpiration, cfg.Cache.DefaultCleanup)
	return internal.NewApi(logger, &cfg, cache, db), nil
}

func database(cfg *config.Config) (*pgxpool.Pool, string, error) {
	dbUriSafe := fmt.Sprintf("postgres://%s:xxxxxxxxxxx@%s/%s?sslmode=disable", cfg.Database.Username, cfg.Database.Server, cfg.Database.Name)

	//TODO: Change this to cfg.Database.Uri()
	dbURI := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", cfg.Database.Username, strings.TrimSpace(cfg.Database.Password), cfg.Database.Server, cfg.Database.Name)

	pool, err := pgxpool.New(context.Background(), dbURI)

	return pool, dbUriSafe, err
}

func explainConfig(cfg *config.Config) {
	fmt.Println("-----------------")
	fmt.Println("BBQd Configured Parameters")
	fmt.Println("-----------------")

	fmt.Printf("Config.Port := %d\n", cfg.Port)
	fmt.Printf("Config.Database.Driver := %s\n", cfg.Database.Driver)
	fmt.Printf("Config.Database.Name := %s\n", cfg.Database.Name)
	fmt.Printf("Config.Database.Username := %s\n", cfg.Database.Username)
	fmt.Printf("Config.Database.Server := %s\n", cfg.Database.Server)
	fmt.Printf("Config.Cache.DefaultExpiration := %s\n", cfg.Cache.DefaultExpiration)
	fmt.Printf("Config.Cache.DefaultCleanup := %s\n", cfg.Cache.DefaultCleanup)
	fmt.Println("-----------------")
}
