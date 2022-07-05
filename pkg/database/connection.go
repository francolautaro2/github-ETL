package database

import (
	"log"
	"repoInfo/pkg/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetConnection() *gorm.DB {
	Instance, err := gorm.Open(mysql.Open("your user database and password here :)"), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}

	log.Println("Connected to databases...!")
	return Instance
}

func Migration() {
	In := GetConnection()
	In.AutoMigrate(&models.Info{})
	log.Println("Migration is succesfully!")
}
