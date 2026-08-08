package main

import (
	"log"

	"go-user-auth/internal/config"
	"go-user-auth/internal/database"
)

func main() {
	config := config.LoadConfig()
	_, err := database.Connect(config)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
}
