package main

import (
	"log"

	"clothes-shop-backend/config"
)

func main() {
	_, err := config.Load()
	if err != nil {
		log.Fatalf("error while config.Load: %v", err)
	}
}
