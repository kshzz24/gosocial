package main

import (
	"log"

	"github.com/kshzz24/gosocial/internal/database"
)

func main() {

	err := database.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.Close()

	log.Println("🚀 Server is ready!")

	// Keep program running (for now, just to test)

}
