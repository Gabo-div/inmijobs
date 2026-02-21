package main

import (
	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/database"
	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/model"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	godotenv.Load()

	db, err := database.NewDatabase()
	if err != nil {
		log.Fatalf("Fatal Error connecting to database: %v", err)
	}

	// Sembrar la reacción por defecto para que las interacciones funcionen
	var reaction model.Reaction
	db.FirstOrCreate(&reaction, model.Reaction{
		Name:    "Me gusta",
		IconURL: "like.png",
	})
	log.Println("INFO [Database] Seeded reactions")

	database.Seed(db)
}
