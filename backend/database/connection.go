package database

import (
	"backend/models"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "root:root@tcp(127.0.0.1:3306)/pesantren?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database", err)
	}
	log.Println("Database connection established")
}

func Migration() {
	ConnectDB()
	err := DB.AutoMigrate(&models.Santri{})
	if err != nil {
		fmt.Println("tdak bisa migrate santris")
	}
	fmt.Println("santris berhasil migrate")
}
