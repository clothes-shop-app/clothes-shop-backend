package main

import (
	handlers "clothes-shop-backend/cmd/internal/handlers/factory"
	repositories "clothes-shop-backend/cmd/internal/repositories/factory"
	services "clothes-shop-backend/cmd/internal/services/factory"
	"clothes-shop-backend/cmd/internal/transport/http"
	"clothes-shop-backend/config"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("error while config.Load: %v", err)
	}

	if cfg.DBURL == "" {
		log.Fatal("DB_URL is required")
	}

	// Remove mysql:// prefix if present for the standard MySQL driver
	dbURL := cfg.DBURL
	log.Printf("Original DB_URL: %s", cfg.DBURL)

	if strings.HasPrefix(dbURL, "mysql://") {
		dbURL = dbURL[len("mysql://"):]
	} else if strings.HasPrefix(dbURL, "mysql:") {
		dbURL = dbURL[len("mysql:"):]
	}
	dbURL = strings.TrimLeft(dbURL, "/")
	log.Printf("Final DB_URL for connection: %s", dbURL)

	db, err := sqlx.Open("mysql", dbURL)
	if err != nil {
		log.Fatalf("error while sql.Open: %v", err)
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		log.Fatalf("error while db.Ping: %v", err)
	}
	log.Println("Successfully connected to database")

	repositories := repositories.NewRepositories(db)
	services := services.NewServices(repositories)
	handlers := handlers.NewHandlers(services)

	http.InitServer(handlers)
}
