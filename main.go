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

// Index
func returnAllApplications(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllApplications")
	json.NewEncoder(w).Encode(Applications)
}

// Show
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

// Create
func createNewApplication(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// return the string response containing the request body
	reqBody, _ := ioutil.ReadAll(r.Body)

	var application Application
	json.Unmarshal(reqBody, &application)
	// update global Applications array
	Applications = append(Applications, application)

	json.NewEncoder(w).Encode(application)
}

// Delete
func deleteApplication(w http.ResponseWriter, r *http.Request) {
	// parse parameters
	vars := mux.Vars(r)
	// get ID of application to delete
	id := vars["id"]

	// loop through applications to find the correct one
	for index, application := range Applications {
		if application.Id == id {
			// if applicaton is found, update Applications array to remove Application
			Applications = append(Applications[:index], Applications[index+1:]...)
			// return deleted application
			json.NewEncoder(w).Encode(application)
			break
		}
	}
}

// Routes, using mux
// they point to functions defined above
func handleRequests() {
	fmt.Println("HTTP Server Started.")

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homePage)
	router.HandleFunc("/applications", returnAllApplications)
	router.HandleFunc("/applications/{id}", returnSingleApplication)
	router.HandleFunc("/applications", createNewApplication).Methods("POST")
	router.HandleFunc("/applications/{id}", deleteApplication).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":10000", router))
}

// this is the Application "model"
type Application struct {
	Id       string `json:"Id"`
	Date     string `json:"Date"`
	Company  string `json:"Company"`
	Position string `json:"Position"`
	Location string `json:"Location"`
}

// global array in place of a database for now
var Applications []Application

func main() {
	Applications = []Application{
		Application{Id: "1", Date: "6/02/20", Company: "Google", Position: "Software Engineer", Location: "Los Angeles, CA"},
		Application{Id: "2", Date: "6/03/20", Company: "Amazon", Position: "Software Engineer", Location: "Los Angeles, CA"},
		Application{Id: "3", Date: "6/05/20", Company: "Dollar Shave Club", Position: "Software Engineer I", Location: "Marina Del Rey, CA"},
	}
	handleRequests()
}
