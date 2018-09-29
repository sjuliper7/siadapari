package controllers

import (
	"POS/thinkservice/models"
	"POS/thinkservice/repository"
)

var think repository.Think

func FetchAllThink() (thinks []models.Think) {
	thinks = think.Fetch()
	return
}
