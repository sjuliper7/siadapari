package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/ziutek/mymysql/godrv"
)

var Db *sql.DB

func OpenConnection() {
	var err error
	Db, err = sql.Open("mymysql", "think_pos/root/password")
	if err != nil {
		log.Panic(err)
	} else {
		fmt.Println("Has been Connected")
	}
}
