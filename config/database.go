package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type dbconfig struct {
	host    string
	port    string
	user    string
	pass    string
	dbname  string
	sslmode string
}

func SetUpDatabase(models ...interface{}) *gorm.DB {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading environment variables: %v", err)
	}

	// set up db credentials
	config := dbconfig{
		host:    os.Getenv("DB_HOST"),
		port:    os.Getenv("DB_PORT"),
		user:    os.Getenv("DB_USER"),
		pass:    os.Getenv("DB_PASS"),
		dbname:  os.Getenv("DB_NAME"),
		sslmode: os.Getenv("DB_SSLMODE"),
	}

	// run db
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v",
		config.host, config.user, config.pass, config.dbname, config.port, config.sslmode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// extension(s) necessary
	err = db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"").Error
	if err != nil {
		log.Fatalf("Error creating extension: %v", err)
	}

	for _, model := range models {
		err = db.AutoMigrate(model)
		if err != nil {
			log.Fatalf("Error migrating model %v in database: %v", model, err)
		}
	}

	return db
}
