package main

import (
	"database/sql"
	"fmt"
)

//NewConnection create connection to mariaDB
func newConnection(u string, p string, db string, host string, port int) (*sql.DB, error) {
	//dbSource := u + ":" + p + "@tcp(" + host + ":" + strconv.Itoa(port) + ")/" + db + "?parseTime=true"
	dbSource := fmt.Sprintf("server=%s;user id=%s;database=%s;password=%s;port=%d", host, u, db, p, port)
	// fmt.Printf("db source: %s", dbSource)
	con, err := sql.Open("mssql", dbSource)
	if err != nil {
		return nil, err
	}

	return con, nil
}

//CloseConnection close connection with mariaDB
func closeConnection(con *sql.DB) error {
	con.Close()
	return nil
}

func getTrackingByUserID(t string, con *sql.DB) ([]Tracking, error) {

	err := con.Ping()
	if err != nil {
		return nil, err
	}

	var ts []Tracking
	rows, err := con.Query(fmt.Sprintf("exec [dbo].[Select_TrackingByUserID] ?", t))
	for rows.Next() {
		var t Tracking
		err = rows.Scan(&t.TrackingID, &t.UserID, &t.SessionID, &t.Stage, &t.Msg, &t.LogTime, &t.QuestionID, &t.Answer)
		ts = append(ts, t)
	}

	defer rows.Close()

	return ts, err
}

func getQuestionByQuestionGroupID(t string, con *sql.DB) ([]Question, error) {

	err := con.Ping()
	if err != nil {
		return nil, err
	}

	var quizs []Question
	rows, err := con.Query("exec [dbo].[Select_QuestionByQuestionGroup] '" + t + "'")
	for rows.Next() {
		var qiz Question
		err = rows.Scan(&qiz.ID, &qiz.QuestionType, &qiz.QuestionTitle, &qiz.QuestionDescription, &qiz.Answer, &qiz.QuestionGroup)
		quizs = append(quizs, qiz)
	}

	defer rows.Close()

	return quizs, err
}

func insertQuestion(q Question, con *sql.DB) (sql.Result, error) {

	err := con.Ping()
	if err != nil {
		return nil, err
	}

	r, err := con.Exec(fmt.Sprintf("exec [dbo].[Insert_Question] ?,?,?,?,?"), q.QuestionType, q.QuestionTitle, q.QuestionDescription, q.Answer, q.QuestionGroup)

	return r, err
}

func insertTracking(t Tracking, con *sql.DB) (sql.Result, error) {

	err := con.Ping()
	if err != nil {
		return nil, err
	}

	r, err := con.Exec(fmt.Sprintf("exec [dbo].[Insert_Tracking] ?,?,?,?,?,?"), t.UserID, t.SessionID, t.Stage, t.Msg, t.QuestionID, t.Answer)

	return r, err
}
