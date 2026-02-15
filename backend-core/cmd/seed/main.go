package main

import (
	"log"

	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/database"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	db, err := database.NewDatabase()
	if err != nil {
		log.Fatalf("Fatal Error connecting to database: %v", err)
	}

	// Run Seed
	database.Seed(db)
}
