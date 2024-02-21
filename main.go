package main

import (
	"log"

	"github.com/adieos/student-db/config"
	"github.com/adieos/student-db/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	server := gin.Default()

	var db *gorm.DB = config.SetUpDatabase(&model.Major{}, &model.Student{})
	var address string = config.SetAddress()

	if err := server.Run(address); err != nil {
		log.Fatalf("Error running server: %v", err)
	}

	log.Printf("server is running %v", db) // JUST so golang thinks db is used lol !!!!!!!!!!!
}
