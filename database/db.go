package database

import (
	"gorm.io/gorm"
	"log"
)
import "gorm.io/driver/postgres"

var (
	DB  *gorm.DB
	err error
)

func Connect() {
	connectionString := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"

	log.Println("Connecting to the database")
	DB, err = gorm.Open(postgres.Open(connectionString))
	if err != nil {
		log.Panic("Failed to connect to the database!")
	}
}
