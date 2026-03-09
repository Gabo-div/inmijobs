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

	// Sembrar la reacción por defecto para que las interacciones funcionen

	log.Println("INFO [Database] Seeded reactions")

	// Run Seed
	database.Seed(db)
}
