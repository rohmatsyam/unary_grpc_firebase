package helpers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func GetEnv() string {
	// projectID := os.Getenv("PROJECT_ID")
	colName := os.Getenv("COL_NAME")

	return colName
}
