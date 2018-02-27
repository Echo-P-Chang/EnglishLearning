package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gorilla/mux"
)

var (
	c      *Config
	sqlcon *sql.DB
	err    error
)

type Question struct {
	ID                  int
	QuestionType        string
	QuestionTitle       string
	QuestionDescription *string
	Answer              string
	QuestionGroup       int
}
type Tracking struct {
	TrackingID int
	UserID     int
	SessionID  string
	Stage      string
	Msg        *string
	LogTime    time.Time
	QuestionID int
	Answer     string
}

func init() {
	flag.Parse()
}

// the main function
func main() {

	c, err = readConfig()
	if err != nil {
		fmt.Println("[Error] error: ", err)
		os.Exit(1)
	}
	sqlcon, err = newConnection(c.DBUser, c.DBPasswd, c.DBDatabase, c.DBHost, c.DBPort)
	if err != nil {
		os.Exit(1)
	}
	err = sqlcon.Ping()
	if err != nil {
		os.Exit(1)
	}

	router := mux.NewRouter()
	router.HandleFunc("/question/{groupid}", GetQuestions).Methods("GET")
	router.HandleFunc("/tracking/{userid}", GetTrackings).Methods("GET")
	router.HandleFunc("/question", CreateQuestion).Methods("POST")
	router.HandleFunc("/tracking", CreateTracking).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))

}

func GetQuestions(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	quizs, _ := getQuestionByQuestionGroupID(params["groupid"], sqlcon)
	json.NewEncoder(w).Encode(quizs)
}

func CreateQuestion(w http.ResponseWriter, r *http.Request) {

	var quiz Question
	json.NewDecoder(r.Body).Decode(&quiz)
	result, err := insertQuestion(quiz, sqlcon)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return
	case nil:
		fmt.Println(result)
	default:
		panic(err)
	}
	json.NewEncoder(w).Encode(quiz)
}

func GetTrackings(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	trackings, _ := getTrackingByUserID(params["userid"], sqlcon)
	json.NewEncoder(w).Encode(trackings)
}

func CreateTracking(w http.ResponseWriter, r *http.Request) {

	var t Tracking
	json.NewDecoder(r.Body).Decode(&t)
	result, err := insertTracking(t, sqlcon)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return
	case nil:
		fmt.Println(result)
	default:
		panic(err)
	}
	json.NewEncoder(w).Encode(t)
}
