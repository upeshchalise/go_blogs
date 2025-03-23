package database

import (
	"log"

	"github.com/spf13/viper"
	"github.com/upeshchalise/go_blogs/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	var err error
	dsn := viper.GetString("DATABASE_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	err = DB.AutoMigrate(&models.User{}, &models.Blog{}, &models.BookMark{}, &models.Comment{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	log.Printf("Database connected successfully")

}
