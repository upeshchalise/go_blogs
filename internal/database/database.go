package database

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
	"github.com/upeshchalise/go_blogs/internal/models"
)

var DB *gorm.DB

func Init() {
	var err error
	dsn := viper.GetString("DATABASE_URL")
	DB, err := gorm.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	DB.AutoMigrate(&models.User{})

	log.Printf("Database connected successfully")

}
