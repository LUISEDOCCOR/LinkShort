package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Conn() {
	var err error
	dsn := os.Getenv("dsn")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Database error 🚨")
	} else {
		fmt.Println("The database is ready to use 🚚")
	}
}
