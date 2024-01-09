package db

import (
	"api-instagram/app/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Instance *gorm.DB

func ConnectDb() {
	DB_USERNAME := "root"
	DB_PASSWORD := "Hyvong1699@"
	DB_HOST := "localhost"
	DB_PORT := "3306"
	DB_NAME := "instagram"
	dsn := DB_USERNAME + ":" + DB_PASSWORD + "@tcp" + "(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?" + "parseTime=true&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db.AutoMigrate(&models.Users{})
	db.AutoMigrate(&models.Posts{})

	if err != nil {
		panic("failed to connect database")
	}
	Instance = db
}
