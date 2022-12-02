package server

import (
	"dyslexia/conf"
    "dyslexia/model"
	"dyslexia/repository"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func getQuestions(w http.ResponseWriter, r *http.Request) {
    results, err := repository.GetQuestions(10)
    if err != nil {
        log.Panic(err)
    }

    json.NewEncoder(w).Encode(results)
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


func calculateEffectiveness(w http.ResponseWriter, r *http.Request) {
    results, err := repository.GetResults()
    if err != nil {
        log.Panic(err)
    }

    questionsEffectiveness := make(map[string]model.Effectiveness)
    var correctIncr int

    for _, result := range results {
        fmt.Println(result)
        for _, question := range result.Responses {
            correctIncr = 0
            if result.Confirmation == question.Status {
                correctIncr = 1
            }

            if effectiveness, ok := questionsEffectiveness[question.QuestionID]; ok {
                effectiveness.Correct += correctIncr
                effectiveness.Total += 1
                questionsEffectiveness[question.QuestionID] = effectiveness
                continue
            }
            questionsEffectiveness[question.QuestionID] = model.Effectiveness{
                Total: 1,
                Correct: correctIncr,
            }
        }
    }

    for id, question := range questionsEffectiveness {
        question.Effectiveness = float64(question.Correct) / float64(question.Total)
        questionsEffectiveness[id] = question
    }

    err = repository.UpdateEffectiveness(questionsEffectiveness)
    if err != nil {
        log.Panic(err)
    }
    json.NewEncoder(w).Encode(questionsEffectiveness)
}

func Start() {
	config := conf.Get()

	r := mux.NewRouter().StrictSlash(true)

    // API Content
	r.HandleFunc("/questions", getQuestions).Methods("GET")
	r.HandleFunc("/sus", getSus).Methods("GET")
	r.HandleFunc("/hospitals", getHospitals).Methods("GET")
	r.HandleFunc("/save-test", getHospitals).Methods("POST")
	r.HandleFunc("/effectiveness", calculateEffectiveness).Methods("GET")

    // Web Content
    r.PathPrefix("/").Handler(http.FileServer(http.Dir("./web/")))

    fmt.Printf("Listen in %s\n", config.Server.Port)
	log.Fatal(http.ListenAndServe(config.Server.Port, r))
}
