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
	"strconv"
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
	XRequestID  string
	SessionID   string
	UserID     int
	LessionID  string
	Event      string
	Note        *string
	LogUTCTime    time.Time
}

type ComboResult struct {
	ComboRecords []ComboRecord
}

type ComboRecord struct {
	UserName string
	UserImgPath string
	UserCombo int
}

type Textbook struct {
	Words []Word
}

type Word struct {
	ID int
	SemesterID int
	LessionID  int
	ImagePath string
	AudioPath string
	Answer      string
	Options        string
}

type Practice struct {

	ID int
	SemesterID int
	LessionID  int
	ImagePath string
	AudioPath string
	Answer      string
	Options        string
	CorrectCnt float32
	CorrectRate float32
	ComboCnt int
}

type PracticeResult struct {
	CorrectCnt float32
	CorrectRate float32
	ComboCnt int

}
type PracticeRecord struct {
	UserPracticeID int
	UserID int
	SessionID string
	SemesterID int
	UserAnswer string
	CorrectAnswer string
	CorrectCnt float32
	CorrectRate float32
	ComboInDB int
	
}
	

type Progress struct {
	ProgressID 	int
	UserID 		int
	SessionID	string
	LessionID	int
}

type CurrentProgress  struct {
 Lessions []int `json: weather`
 }
 
type User struct {
	UserID          int
	UserName        string
	PhotoPath       string	
	LocalPhotoPath       string

}

type AnswerSet struct {
	Answers []Answer
}
type Answer struct {
	AnswerID	int
    UserID		int
    SessionID	string
    SemesterID	int
    LessionID	int
    UserAnswer	string
    CorrectAnswer string
}

type StairResult struct {
	StairRecords []StairRecord
}
type StairRecord struct {
	UserName string
	LocalPhotoPath string
	TotalCnt int
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
		fmt.Println("[Error] error: ", err)
		os.Exit(1)
	}
	err = sqlcon.Ping()
	if err != nil {
		fmt.Println("[Error] error: ", err)
		os.Exit(1)
	}

	router := mux.NewRouter()
	router.HandleFunc("/question/{groupid}", GetQuestions).Methods("GET")
	router.HandleFunc("/user/{userid}", GetUserInfo).Methods("GET")
	router.HandleFunc("/combo/{sessionid}", GetComboDataBySessionID).Methods("GET")	
	router.HandleFunc("/stair/{sessionid}", GetStairResultBySessionID).Methods("GET")
	router.HandleFunc("/word/{id}", GetOneWord).Methods("GET")

	

	
	//router.HandleFunc("/tracking/{lessionID}/{deviceID}", GetTrackingsByDeviceIDLessionID).Methods("GET")
	//router.Queries("deviceID")
	
	router.HandleFunc("/tracking/{userid}/{sessionid}", GetTrackingsByUserIDSessionID).Methods("GET")
	router.Queries("sessionid")
	
	router.Handle("/photo/{userid}", http.StripPrefix("/photo/", http.FileServer(http.Dir("photo"))))
    
	router.Handle("/texture/{semester}/{lession}/{name}", http.StripPrefix("/texture/", http.FileServer(http.Dir("texture"))))
    
	
	router.HandleFunc("/question", CreateQuestion).Methods("POST")
	router.HandleFunc("/tracking", CreateTracking).Methods("POST")
	router.HandleFunc("/answer", UpsertAnswer).Methods("POST")
	router.HandleFunc("/practice", CreateUserPracticeRecord).Methods("POST")
	
	router.HandleFunc("/progress", CreateProgress).Methods("POST")
	router.HandleFunc("/progress/{userid}/{sessionid}", GetProgressByUserIDSessionID).Methods("GET")
	router.Queries("sessionid")
	
	//words
	router.HandleFunc("/words/{semesterid}", GetWordsBySemesterID).Methods("GET")
	router.HandleFunc("/word", CreateWord).Methods("POST")
	
	router.HandleFunc("/user", CreateUser).Methods("POST")
	
	//answer	
	router.HandleFunc("/answer/{userid}/{sessionid}", GetAnswerByUserIDSessionID).Methods("GET")
	router.Queries("sessionid")

	//practice
	
	router.HandleFunc("/practice/{userid}/{sessionid}", GetPracticeByUserIDSessionID).Methods("GET")
	router.Queries("sessionid")
	
	log.Fatal(http.ListenAndServe(":8000", router))

}

func GetQuestions(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	quizs, _ := getQuestionByQuestionGroupID(params["groupid"], sqlcon)
	json.NewEncoder(w).Encode(quizs)
}

func GetComboDataBySessionID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	
    output := ComboResult{}	
	result, _ := getComboFromDBBySessionID(params["sessionid"], sqlcon)

	output.ComboRecords = result
	
	json2, _ := json.Marshal(output)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json2)
	
}
func GetStairResultBySessionID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	
    output := StairResult{}	
	result, _ := GetStairRecordsFromDBBySessionID(params["sessionid"], sqlcon)

	output.StairRecords = result
	
	json2, _ := json.Marshal(output)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json2)
	
}
func GetOneWord(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	wid,_:=strconv.Atoi(params["id"]) 

	words, _ := getOneWordByWordID(wid, sqlcon)
	
	
    output := Textbook{}
	output.Words = words
	
	json2, _ := json.Marshal(output)

	
	//output := CurrentProgress{}
	//output.Lessions = progress
	
	//json2, _ := json.Marshal(output)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json2)
	
	//json.NewEncoder(w).Encode(json2)
	
}
 


func GetWordsBySemesterID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	words, _ := getWordsBySemesterIDFromDB(params["semesterid"], sqlcon)
	
	
    output := Textbook{}
	output.Words = words
	
	json2, _ := json.Marshal(output)

	
	//output := CurrentProgress{}
	//output.Lessions = progress
	
	//json2, _ := json.Marshal(output)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json2)
	
	//json.NewEncoder(w).Encode(json2)
}

func GetUserInfo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userInfo, _ := getUserInfoByUserID(params["userid"], sqlcon)
	json.NewEncoder(w).Encode(userInfo)
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


func CreateUser(w http.ResponseWriter, r *http.Request) {

	var usrs []User
	json.NewDecoder(r.Body).Decode(&usrs)
	for _, v := range usrs{
		result, err := insertOneUser(v, sqlcon)
		switch err {
		case sql.ErrNoRows:
			fmt.Println("No rows were returned!")
			return
		case nil:
			fmt.Println(result)
		default:
			panic(err)
		}
	}
//	json.NewEncoder(w).Encode(usrs)
}

func GetTrackings(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	trackings, _ := getTrackingByUserID(params["userid"], sqlcon)
	json.NewEncoder(w).Encode(trackings)
}
func CreateProgress(w http.ResponseWriter, r *http.Request) {

	var p Progress
	json.NewDecoder(r.Body).Decode(&p)
	result, err := insertProgress(p, sqlcon)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return
	case nil:
		fmt.Println(result)
	default:
		panic(err)
	}
	json.NewEncoder(w).Encode(p)
}

func CreateUserPracticeRecord(w http.ResponseWriter, r *http.Request) {

	var p PracticeRecord
	json.NewDecoder(r.Body).Decode(&p)
	result, err := insertUserPractice(p, sqlcon)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return
	case nil:
		fmt.Println(result)
	default:
		panic(err)
	}
	json.NewEncoder(w).Encode(p)
}

func CreateWord(w http.ResponseWriter, r *http.Request) {

	var wds []Word
	json.NewDecoder(r.Body).Decode(&wds)
	
	for _, v := range wds{
		result, err := insertWord(v, sqlcon)
		switch err {
		case sql.ErrNoRows:
			fmt.Println("No rows were returned!")
			return
		case nil:
			fmt.Println(result)
		default:
			panic(err)
		}
	}
	
	json.NewEncoder(w).Encode(wds)
}

func GetTrackingsByDeviceIDLessionID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	trackings, _ := getTrackingByDeviceIDLessionID(params["deviceID"], params["lessionID"], sqlcon)
	json.NewEncoder(w).Encode(trackings)
}

func GetTrackingsByUserIDSessionID(w http.ResponseWriter, r *http.Request) {

	//fmt.Println("No rows were returned!")
	params := mux.Vars(r)
	//uid,_:=strconv.Atoi(params["userid"]) 
	//fmt.Println(uid)
	//fmt.Println( params["sessionid"])
	
	trackings, _ := getTrackingByUserIDSessionID(params["userid"], params["sessionid"], sqlcon)
	json.NewEncoder(w).Encode(trackings)
}

func GetProgressByUserIDSessionID(w http.ResponseWriter, r *http.Request) {

	//fmt.Println("No rows were returned!")
	params := mux.Vars(r)
	uid,_:=strconv.Atoi(params["userid"]) 
	//fmt.Println(uid)
	
	progress, _ := getProgressesByUserIDSessionID(uid, params["sessionid"], sqlcon)
	
	output := CurrentProgress{}
	output.Lessions = progress
	
	json2, _ := json.Marshal(output)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json2)
}

func GetAnswerByUserIDSessionID(w http.ResponseWriter, r *http.Request) {

	//fmt.Println("No rows were returned!")
	params := mux.Vars(r)
	uid,_:=strconv.Atoi(params["userid"]) 
	//fmt.Println(uid)
	
	as, _ := getAnswerByUserIDSessionIDFromDB(uid, params["sessionid"], sqlcon)
	
	output := AnswerSet{}
	output.Answers = as
	
	json2, _ := json.Marshal(output)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json2)
}
func GetPracticeByUserIDSessionID(w http.ResponseWriter, r *http.Request) {

	//fmt.Println("No rows were returned!")
	params := mux.Vars(r)
	uid,_:=strconv.Atoi(params["userid"]) 
	//fmt.Println(uid)
	
	wd, _ := getPracticeByUserIDSessionIDFromDB(uid, params["sessionid"], sqlcon)
	
	result, _ := getUserPracticeResultByUserIDSessionID(uid, params["sessionid"], sqlcon)
	
	
	pp := Practice{}
	pp.ID = wd.ID
	pp.SemesterID = wd.SemesterID
	pp.LessionID = wd.LessionID
	pp.ImagePath = wd.ImagePath
	pp.AudioPath = wd.AudioPath
	pp.Answer = wd.Answer
	pp.Options = wd.Options
	

	pp.CorrectCnt = result.CorrectCnt
	pp.CorrectRate = result.CorrectRate
	pp.ComboCnt = result.ComboCnt
	
	fmt.Println(pp.AudioPath)
	
	json2, _ := json.Marshal(pp)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json2)
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

func UpsertAnswer(w http.ResponseWriter, r *http.Request) {

	var a Answer
	json.NewDecoder(r.Body).Decode(&a)
	result, err := upsertAnswerToDB(a, sqlcon)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return
	case nil:
		fmt.Println(result)
	default:
		panic(err)
	}
	json.NewEncoder(w).Encode(a)
}
