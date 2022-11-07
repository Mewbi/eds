package server

import (
	"dyslexia/conf"
	"dyslexia/repository"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

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
    results, err := repository.GetResults()
    if err != nil {
        log.Panic(err)
    }
	json.NewEncoder(w).Encode(results)
}

func Start() {
	config := conf.Get()

	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/questions", getQuestions).Methods("GET")
	r.HandleFunc("/sus", getSus).Methods("GET")
	r.HandleFunc("/hospitals", getHospitals).Methods("GET")
	r.HandleFunc("/save-test", getHospitals).Methods("POST")
	r.HandleFunc("/efficacy", calculateEfficacy).Methods("GET")

    fmt.Printf("Listen in %s\n", config.Server.Port)
	log.Fatal(http.ListenAndServe(config.Server.Port, r))
}
