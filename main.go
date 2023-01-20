package main

import (
	"log"
	"rest/gin/Config"
	"rest/gin/models"

	"rest/gin/Routes"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var err error

func init() {
	Config.DB, err = gorm.Open(mysql.Open(Config.DbURL(Config.BuildDBConfig())), &gorm.Config{})
	if err != nil {
		log.Fatal("Status:", err)
	}

	Config.DB.AutoMigrate(&models.Employee{})

}

func main() {

	r := Routes.SetupRouter()
	r.Run()
}
