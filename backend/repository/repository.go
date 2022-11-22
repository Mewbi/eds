package repository

import (
	"database/sql"
	"dyslexia/conf"
	"dyslexia/model"
	"encoding/json"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Question struct {
    ID string `json:"question_id"`
    Content string `json:"content"`
    Count int `json:"count"`
}

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

func GetQuestions(numberQuestions int) ([]Question, error) {
    var questions []Question
	config := conf.Get()
	db, err := sql.Open("sqlite3", config.Db.Address)
    if err != nil {
        return questions, err
    }
    defer db.Close()
    rows, err := db.Query(fmt.Sprintf(`SELECT q.id, q.content, qe.total FROM questions AS q INNER JOIN questions_effectiveness AS qe ON q.id = qe.question_id ORDER BY qe.total DESC LIMIT %d`, numberQuestions))
    if err != nil {
        return questions, err
    }

    for rows.Next() {
        var id string
        var content string
        var total int

        err = rows.Scan(&id, &content, &total)
        if err != nil {
            continue
        }
        questions = append(questions, Question{
            ID: id,
            Content: content,
            Count: total,
        })
    }
	return questions, nil
}

func GetResults() ([]Result, error ){
    var results []Result
	config := conf.Get()
	db, err := sql.Open("sqlite3", config.Db.Address)
    if err != nil {
        return results, err
    }
    defer db.Close()
    rows, err := db.Query(`SELECT * FROM test_results WHERE confirmation IS NOT NULL`)
    if err != nil {
        return results, err
    }

    for rows.Next() {
        var id string
        var name string
        var email string
        var responsesString string
        var createdAt time.Time
        var confirmation bool
        var responses []Response

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

func UpdateEffectiveness(questionsEffectiveness map[string]model.Effectiveness) error {
	config := conf.Get()
	db, err := sql.Open("sqlite3", config.Db.Address)
    if err != nil {
        return err
    }
    defer db.Close()
    query, err := db.Prepare(`DELETE FROM questions_effectiveness`)
    if err != nil {
        return err
    }

    _, err = query.Exec()
    if err != nil {
        return err
    }

    query, err = db.Prepare(`INSERT INTO questions_effectiveness (question_id, total, effectiveness) VALUES (?, ?, ?)`)
    if err != nil {
        return err
    }

    for id, effectiveness := range questionsEffectiveness {
        _, err = query.Exec(id, effectiveness.Total, effectiveness.Effectiveness)
        if err != nil {
            return err
        }
    }
    return nil
}
