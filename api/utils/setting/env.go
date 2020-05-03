package setting

import (
	"log"
	"github.com/joho/godotenv"
)

// Setup initialize the configuration instance
func Setup() {
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}
}

