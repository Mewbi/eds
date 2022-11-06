package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const port string = ":8000" 

func getQuestions(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode("Returning question")
}

func getSus(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode("Returning sus")
}

func getHospitals(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode("Returning hospitals")
}

func saveTest(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode("Saving test results in database xD")
}

func calculateEfficacy(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode("Returning effficacy of questions")
}

func handleRequests() {
    r := mux.NewRouter().StrictSlash(true)
    
    r.HandleFunc("/questions", getQuestions).Methods("GET")
    r.HandleFunc("/sus", getSus).Methods("GET")
    r.HandleFunc("/hospitals", getHospitals).Methods("GET")
    r.HandleFunc("/save-test", getHospitals).Methods("POST")
    r.HandleFunc("/efficay", getHospitals).Methods("GET")
    
    log.Fatal(http.ListenAndServe(port, r))
}


func main() {
    fmt.Println("Listen")
    handleRequests()
}
