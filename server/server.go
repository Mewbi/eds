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
    "github.com/gorilla/schema"
)

func getQuestions(w http.ResponseWriter, r *http.Request) {
    results, err := repository.GetQuestions(10)
    if err != nil {
        log.Panic(err)
    }

    json.NewEncoder(w).Encode(results)
}

func createComment(w http.ResponseWriter, r *http.Request) {
    var comment model.Comment

    err := r.ParseForm()
    if err != nil {
        log.Panicf("Error reading body %v", err)
    }

    decoder := schema.NewDecoder()
    err = decoder.Decode(&comment, r.Form)
    if err != nil {
         log.Panicf("Error decoding form: %v", err)
    }

    err = repository.CreateComment(comment.Name, comment.Comment)
    if err != nil {
         log.Panicf("Error creating comment: %v", err)
    }

	json.NewEncoder(w).Encode("Success")
}

func getComments(w http.ResponseWriter, r *http.Request) {
    var auth model.Auth
    var comments []model.Comment
    config := conf.Get()

    err := r.ParseForm()
    if err != nil {
        log.Panicf("Error reading body %v", err)
    }

    decoder := schema.NewDecoder()
    err = decoder.Decode(&auth, r.Form)
    if err != nil {
        log.Panicf("Error decoding form: %v", err)
    }

    if auth.Auth != config.Server.Auth {
        json.NewEncoder(w).Encode(comments)
        return
    }

    comments, err = repository.GetComments()
    if err != nil {
        log.Panicf("Error getting comments: %v", err)
    }

    json.NewEncoder(w).Encode(comments)
}

func getSus(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Returning sus")
}

func getHospitals(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Returning hospitals")
}

func saveTest(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Success")
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
	r.HandleFunc("/comment", createComment).Methods("POST")
	r.HandleFunc("/comment/view", getComments).Methods("POST")
	r.HandleFunc("/effectiveness", calculateEffectiveness).Methods("GET")

    // Web Content
    r.PathPrefix("/").Handler(http.FileServer(http.Dir("./web/")))

    fmt.Printf("Listen in %s\n", config.Server.Port)
	log.Fatal(http.ListenAndServe(config.Server.Port, r))
}
