package main

import (
	"log"
	"rest/gin/Config"
	"rest/gin/Models"
	"rest/gin/Routes"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var err error
var dbDriver string = "mysql"

func init() {
	Config.DB, err = gorm.Open(mysql.Open(Config.DbURL(Config.BuildDBConfig())), &gorm.Config{})
	if err != nil {
		log.Fatal("Status:", err)
	}

	Config.DB.AutoMigrate(&Models.Employee{})
}

func main() {

	r := Routes.SetupRouter()
	r.Run()
}
