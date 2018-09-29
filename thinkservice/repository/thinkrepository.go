package repository

import (
	Database "POS/thinkservice/db"
	"POS/thinkservice/models"
	"log"
)

type Think models.Think

type ThinkRepository interface {
	Fetch() (thinks []models.Think)
}

func checkErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func (t Think) Fetch() (thinks []models.Think) {

	query := "SELECT * FROM think"
	rows, err := Database.Db.Query(query)
	checkErr(err)
	for rows.Next() {
		var think models.Think
		err = rows.Scan(&think.Id, &think.Name, &think.Price, &think.Quantity, &think.Description)
		checkErr(err)
		thinks = append(thinks, think)
	}

	return
}
