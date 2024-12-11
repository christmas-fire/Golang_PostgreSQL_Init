package main

import (
	"log"

	"github.com/christmas-fire/Golang_PostgreSQL_Init/internal/database"
)

func main() {
	db := database.InitDB()
	defer db.Close()

	err := database.CreateTable(db)
	if err != nil {
		log.Fatal(err)
	}

}
