package webserver

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type someResponse struct {
	AccountId int
	Name      string
}

// Initialize our routes
var routes = Routes{

	Route{
		"GetAccount", // Name
		"GET",        // HTTP method
		"/accounts/{accountId}", // Route pattern
		func(w http.ResponseWriter, r *http.Request) {
			log.Println("Got request from " + r.RemoteAddr)
			vars := mux.Vars(r)
			accountId, err := strconv.Atoi(vars["accountId"])
			if err != nil {
				log.Fatalf("Cannot parse %v, error was %v", vars["accountId"], err)
				w.WriteHeader(http.StatusBadRequest)
			} else {
				w.Header().Set("Content-Type", "application/json; charset=UTF-8")
				resD := &someResponse{AccountId: accountId, Name: "Hans Wurst"}
				resB, _ := json.Marshal(resD)
				w.Write(resB)
			}
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
	r := NewRouter(routes)                    // NEW
	http.Handle("/", r)                       // NEW
	err := http.ListenAndServe(":"+port, nil) // Goroutine will block here
	if err != nil {
		log.Fatal("An error occured starting HTTP listener at port " + port)
		log.Fatal("Error: " + err.Error())
	}
}
