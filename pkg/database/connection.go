package database

import (
	"log"
	"repoInfo/pkg/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetConnection() *gorm.DB {
	Instance, err := gorm.Open(mysql.Open("root:admin@tcp(localhost:3306)/gitapi"), &gorm.Config{})
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
