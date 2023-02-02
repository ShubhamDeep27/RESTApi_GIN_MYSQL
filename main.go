package main

import (
	"log"
	"rest/gin/config"
	"rest/gin/models"

	"rest/gin/routes"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var err error

func init() {
	config.DB, err = gorm.Open(mysql.Open(config.DbURL()), &gorm.Config{})
	if err != nil {
		log.Fatal("Status:", err)
	}

	config.DB.AutoMigrate(&models.Employee{})

}

func main() {

	r := routes.SetupRouter()
	r.Run()
}
