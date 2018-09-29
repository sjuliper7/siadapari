package service

import (
	"POS/thinkservice/controllers"
	"encoding/json"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Defines the type Routes which is just an array (slice) of Route structs.
type Routes []Route

// Initialize our routes
var routes = Routes{

	Route{
		"GetAccount", // Name
		"GET",        // HTTP method
		"/index",     // Route pattern
		index,
	},
	Route{
		"Test",
		"GET",
		"/test",
		test,
	},
}

type Response struct {
	Message string `json:"message"`
}

func index(w http.ResponseWriter, r *http.Request) {
	response := Response{Message: "Hello World i'm micro"}
	data, err := json.Marshal(response)

	if err != nil {
		panic("Opps")
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write([]byte(data))

}

func test(w http.ResponseWriter, r *http.Request) {
	data, err := json.Marshal(controllers.FetchAllThink())

	if err != nil {
		panic("Opps")
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write([]byte(data))

}
