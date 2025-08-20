package main

import (
	"log"

	"yourproject/configs"
	"yourproject/internal/db"
	"yourproject/internal/http"
)

func main() {
	// Load configuration
	cfg := configs.LoadConfig()

	// Connect to database
	database, err := db.ConnectDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer func() {
		if err := db.CloseDB(database); err != nil {
			log.Printf("Error closing database connection: %v", err)
		}
	}()

	// Start the server
	log.Printf("Starting user management API server on port %s", cfg.ServerPort)
	if err := http.StartServer(cfg, database); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
