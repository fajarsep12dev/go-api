package setting

import (
	"log"
	"github.com/joho/godotenv"
)

// Initialize the configuration instance
func Initialize() {
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}
}

