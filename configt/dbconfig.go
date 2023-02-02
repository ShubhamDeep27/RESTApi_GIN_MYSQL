package config

import (
	"fmt"
	"rest/gin/util"

	"gorm.io/gorm"
)

var DB *gorm.DB

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

func DbURL() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		util.GetEnvVariable("DB_USER"),
		util.GetEnvVariable("DB_PASSWORD"),
		util.GetEnvVariable("DB_HOST"),
		util.GetEnvVariable("DB_PORT"),
		util.GetEnvVariable("DB_NAME"),
	)
}
