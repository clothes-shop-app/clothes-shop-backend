package main

import (
	handlers "clothes-shop-backend/internal/handlers/factory"
	repositories "clothes-shop-backend/internal/repositories/factory"
	services "clothes-shop-backend/internal/services/factory"
	"clothes-shop-backend/internal/transport/http"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found or error loading .env file:", err)
	}

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
