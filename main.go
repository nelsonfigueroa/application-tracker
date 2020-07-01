package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
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

func returnSingleApplication(w http.ResponseWriter, r *http.Request) {
	// you can extract variables in URLs, and we use that to find the corresponding application
	vars := mux.Vars(r)
	key := vars["id"]

	fmt.Fprintf(w, "Key: "+key)

	// loop through applications to find the correct one
	for _, application := range Applications {
		if application.Id == key {
			json.NewEncoder(w).Encode(application)
		}
	}
}

func createNewApplication(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// return the string response containing the request body
	reqBody, _ := ioutil.ReadAll(r.Body)

	var application Application
	json.Unmarshal(reqBody, &application)
	// update our global Articles array to include
	// our new Article
	Applications = append(Applications, application)

	json.NewEncoder(w).Encode(application)
}

// these are the routes, using mux
func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homePage)
	router.HandleFunc("/applications", returnAllApplications)
	router.HandleFunc("/application/{id}", returnSingleApplication)
	router.HandleFunc("/application", createNewApplication).Methods("POST")

	log.Fatal(http.ListenAndServe(":10000", router))
}

type Application struct {
	Id       string `json:"Id"`
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
		Application{Id: "1", Date: "6/02/20", Company: "Google", Position: "Software Engineer", Location: "Los Angeles, CA"},
		Application{Id: "2", Date: "6/03/20", Company: "Amazon", Position: "Software Engineer", Location: "Los Angeles, CA"},
		Application{Id: "3", Date: "6/05/20", Company: "Dollar Shave Club", Position: "Software Engineer I", Location: "Marina Del Rey, CA"},
	}
	handleRequests()
}
