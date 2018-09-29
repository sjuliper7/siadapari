package service

import (
	"log"
	"net/http"
)

func StartWebServer(port string) {
	log.Println("Starting HTTP Service at " + port)

	//add Router
	r := NewRouter()    // NEW
	http.Handle("/", r) // NEW

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Println("An error occured starting HTTP listener at port " + port)
		log.Println("Error: " + err.Error())
	}
}
