package webserver

import (
	"log"
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
		"/accounts/{accountId}", // Route pattern
		func(w http.ResponseWriter, r *http.Request) {
			log.Println("Got request from " + r.RemoteAddr)
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.Write([]byte("{\"result\":\"OK\"}"))
		},
	},

	Route{
		"GetAlive", // Name
		"GET",      // HTTP method
		"/alive",   // Route pattern
		func(w http.ResponseWriter, r *http.Request) {
			log.Println("Got request from " + r.RemoteAddr)
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.Write([]byte("{\"result\":\"OK\"}"))
		},
	},
}

func StartWebServer(port string) {
	log.Println("Starting HTTP service at " + port)
	r := NewRouter()                          // NEW
	http.Handle("/", r)                       // NEW
	err := http.ListenAndServe(":"+port, nil) // Goroutine will block here
	if err != nil {
		log.Println("An error occured starting HTTP listener at port " + port)
		log.Println("Error: " + err.Error())
	}
}
