package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllApplications(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllApplications")
	json.NewEncoder(w).Encode(Applications)
}

// these are the routes, using mux
func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homePage)
	router.HandleFunc("/applications", returnAllApplications)

	log.Fatal(http.ListenAndServe(":10000", router))
}

type Application struct {
	Date     string `json:"Date"`
	Company  string `json:"Company"`
	Position string `json:"Position"`
	Location string `json:"Location"`
}

// let's declare a global Applications array
// that we can then populate in our main function
// to simulate a database
var Applications []Application

func main() {
	Applications = []Application{
		Application{Date: "6/02/20", Company: "Google", Position: "Software Engineer", Location: "Los Angeles, CA"},
	}
	handleRequests()
}
