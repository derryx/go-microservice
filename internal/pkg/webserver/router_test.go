package webserver

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"testing"

	"github.com/gorilla/mux"
	. "github.com/smartystreets/goconvey/convey"
)

func TestRouting(t *testing.T) {
	// Only pass t into top-level Convey calls
	Convey("Calling NewRouter() should give a new router with routes", t, func() {
		routes = Routes{

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
		router := NewRouter(routes)

		Convey("And it should have some routes", func() {
			route := router.GetRoute("GetAccount")
			So(route, ShouldNotBeNil)
		})
	})
}
