package main

import (
	handlers "clothes-shop-backend/cmd/internal/handlers/factory"
	repositories "clothes-shop-backend/cmd/internal/repositories/factory"
	services "clothes-shop-backend/cmd/internal/services/factory"
	"clothes-shop-backend/cmd/internal/transport/http"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL is required")
	}

	db, err := sqlx.Open("mysql", dbURL)
	if err != nil {
		log.Fatalf("error while sql.Open: %v", err)
	}

	repositories := repositories.NewRepositories(db)
	services := services.NewServices(repositories)
	handlers := handlers.NewHandlers(services)

	http.InitServer(handlers)
}
