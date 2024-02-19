package main

import (
	"log"

	"github.com/gin-gonic/gin"
	// "gorm.io/gorm"
)

func main() {
	server := gin.Default()

	if err := server.Run(":8080"); err != nil {
		log.Fatalf("Error running server: %v", err)
	}

	log.Print("server is running")
}
