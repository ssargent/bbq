package main

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"os"
)

func main() {

	password := os.Getenv("bbq-postgres-password")

	dburl := fmt.Sprintf("postgres://bbqdbuser:%s@192.168.1.40/bbq?sslmode=disable", password)

	m, err := migrate.New(
		"file://migrations",
		dburl)

	if err != nil {
		fmt.Println(err)
	}

	m.Up()
}
