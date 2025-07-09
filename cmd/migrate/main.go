package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const (
	migrationsPath = "file://migrations"
)

func main() {
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL is required")
	}

	action := flag.String("action", "", "Available commands: up, down, force, version, drop")
	version := flag.Int("version", 0, "Migration version number")
	flag.Parse()

	m, err := migrate.New(migrationsPath, dbURL)
	if err != nil {
		log.Fatalf("Error creating migrator: %v", err)
	}

	switch *action {
	case "up":
		err = m.Up()
		if err != nil && err != migrate.ErrNoChange {
			log.Fatalf("Error applying migration up: %v", err)
		}
		fmt.Println("Migrations applied successfully")

	case "down":
		err = m.Down()
		if err != nil && err != migrate.ErrNoChange {
			log.Fatalf("Error applying migration down: %v", err)
		}
		fmt.Println("Migrations rolled back")

	case "force":
		if *version == 0 {
			log.Fatal("Please specify a version number for force")
		}
		err = m.Force(*version)
		if err != nil {
			log.Fatalf("Error applying force: %v", err)
		}
		fmt.Printf("Migration forcibly set to version: %d\n", *version)

	case "version":
		v, dirty, vErr := m.Version()
		if vErr != nil {
			log.Fatalf("Error getting migration version: %v", vErr)
		}
		fmt.Printf("Current migration version: %d (dirty: %v)\n", v, dirty)

	case "drop":
		err = m.Drop()
		if err != nil {
			log.Fatalf("Error dropping all migrations: %v", err)
		}
		fmt.Println("All migrations have been dropped")

	default:
		fmt.Println("Unknown command. Use: up, down, force, version, drop")
	}
}
