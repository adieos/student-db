package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func SetAddress() string {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading environment variables: %v", err)
	}

	host := os.Getenv("APP_HOST")
	port := os.Getenv("APP_PORT")
	address := fmt.Sprintf("%v:%v", host, port)

	return address
}
