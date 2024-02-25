package main

import (
	"flag"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	var storageURL, migrationsPath, migrationsTable string

	flag.StringVar(&storageURL, "storage-url", "", "URL to connect to the database")
	flag.StringVar(&migrationsPath, "migrations-path", "", "Path to migrations")
	flag.StringVar(&migrationsTable, "migrations-table", "migrations", "Name of migrations table")
	flag.Parse()

	validateParams(storageURL, migrationsPath)

	m, err := migrate.New(
		"file://"+migrationsPath,
		storageURL,
	)
	if err != nil {
		panic(err)
	}

	if err := m.Up(); err != nil {
		if err == migrate.ErrNoChange {
			fmt.Println("No migrations to apply")
			return
		}
		panic(err)
	}

	fmt.Println("Migrations applied")
}

func validateParams(storageURL, migrationsPath string) {
	if storageURL == "" {
		panic("storage-url is required")
	}
	if migrationsPath == "" {
		panic("migrations-path is required")
	}
}

// Log represents the logger
type Log struct {
	verbose bool
}

// Printf prints out formatted string into a log
func (l *Log) Printf(format string, v ...interface{}) {
	fmt.Printf(format, v...)
}

// Verbose shows if verbose print enabled
func (l *Log) Verbose() bool {
	return false
}
