package initializers

import (
	"feed-me/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := os.Getenv("db_url")

	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to the database:", err)
	}

	log.Println("Database connected successfully!")
}

func Migrate() {
	DB.AutoMigrate(&models.UserModel{})
}
