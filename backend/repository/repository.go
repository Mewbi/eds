package repository

import (
	"database/sql"
	"dyslexia/conf"
	"encoding/json"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Response struct {
    QuestionID string `json:"question_id"`
    Status bool `json:"response"`
}

type Result struct {
    ID string
    Name string
    Email string
    Responses []Response
    CreatedAt time.Time
    Confirmation bool
}

func GetResults() ([]Result, error ){
    var results []Result
	config := conf.Get()
	db, err := sql.Open("sqlite3", config.Db.Address)
    if err != nil {
        return results, err
    }
    defer db.Close()
    rows, err := db.Query("SELECT * FROM test_results")
    if err != nil {
        return results, err
    }

    var id string
    var name string
    var email string
    var responsesString string
    var createdAt time.Time
    var confirmation bool
    var responses []Response
    for rows.Next() {
        err = rows.Scan(&id, &name, &email, &responsesString, &createdAt, &confirmation)
        if err != nil {
            continue
        }
        json.Unmarshal([]byte(responsesString), &responses)
        results = append(results, Result{
            ID: id,
            Name: name,
            Email: email,
            Responses: responses,
            CreatedAt: createdAt,
            Confirmation: confirmation,
        })
    }
	return results, nil
}
