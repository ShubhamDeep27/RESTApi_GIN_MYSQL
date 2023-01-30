package main

import (
	"database/sql"
	"log"
	Config "rest/gin/config"

	Routes "rest/gin/routes"
)

var err error

func init() {
	Config.DB, err = sql.Open("mysql", Config.DbURL())
	if err != nil {
		log.Fatal("Status:", err)
	}

}

func main() {

	r := Routes.SetupRouter()
	r.Run()
}
