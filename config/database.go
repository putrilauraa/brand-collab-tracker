package config

import (
	"fmt"
	"log"
	"os"

	"brand-collab-tracker/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	LoadEnv()

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		os.Getenv("PGHOST"),
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
		os.Getenv("PGDATABASE"),
		os.Getenv("PGPORT"),
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL database: %v", err)
	}

	log.Println("Successfull database connection!")

	err = DB.AutoMigrate(
		&models.User{},
		&models.CategoryMaster{},
		&models.Brand{},
		&models.Project{},
		&models.Task{},
		&models.ProjectAttachment{},
	)

	if err != nil {
		log.Fatalf("Failed to migrate table: %v", err)
	}
	log.Println("Successfull table migration!")
}