package main

import (
	"POS/thinkservice/db"
	"POS/thinkservice/service"
	"fmt"
)

var appName = "thinkservice"

func main() {

	fmt.Println("Starting", appName)
	initDb()
	service.StartWebServer("6767")

}

func initDb() {
	fmt.Println("Database opened..")
	db.OpenConnection()
}
