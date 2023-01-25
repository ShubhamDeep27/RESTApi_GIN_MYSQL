package main

import (
	"database/sql"
	"log"
	"rest/gin/Config"

	"rest/gin/Routes"
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
