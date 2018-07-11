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
	rows, err := con.Query("exec [dbo].[Select_TrackingByUserID] '" + t + "'")
	for rows.Next() {
		var t Tracking
		err = rows.Scan(&t.TrackingID, &t.SessionID, &t.UserID, &t.LessionID, &t.Event, &t.Note, &t.LogUTCTime)
		ts = append(ts, t)
	}

	defer rows.Close()

	return ts, err
}

func getTrackingByDeviceIDLessionID(deviceID string, lessionID string, con *sql.DB) ([]Tracking, error) {

	err := con.Ping()
	if err != nil {
		return nil, err
	}

	var ts []Tracking
	
	rows, err := con.Query(fmt.Sprintf("exec [dbo].[Select_TrackingByDeviceIDLessionID] ?,?"), deviceID, lessionID)
	for rows.Next() {
		var t Tracking
		err = rows.Scan(&t.TrackingID, &t.SessionID, &t.UserID, &t.LessionID, &t.Event, &t.Note, &t.LogUTCTime)
		ts = append(ts, t)
	}

	defer rows.Close()

	return ts, err
}



func getTrackingByUserIDSessionID(userid string, sessionid string, con *sql.DB) ([]Tracking, error) {

	err := con.Ping()
	if err != nil {
		return nil, err
	}

	var ts []Tracking
	//fmt.Println(fmt.Sprintf("exec [dbo].[Select_TrackingByUserIDSessionID] ?,?"), userid, sessionid)
	//TrackingID	XRequestID	SessionID	UserID	LessionID	Event	Note	LogUTCTime
	//1178	430d3147-4032-4ffb-8c35-48ec2d9a5eed	20180704	105554002	0	USER_LOGGINGIN	User:105554002 Logging in at:0	2018-07-04 06:17:22.313
	//TrackingID int
	//XRequestID  string
	//SessionID   string
	//UserID     int
	//LessionID  string
	//Event      string
	//Note        *string
	//LogUTCTime    time.Time
	
	rows, err := con.Query(fmt.Sprintf("exec [dbo].[Select_TrackingByUserIDSessionID] ?,?"), userid, sessionid)
	for rows.Next() {
		var t Tracking
		err = rows.Scan(&t.TrackingID, &t.XRequestID, &t.SessionID, &t.UserID, &t.LessionID, &t.Event, &t.Note, &t.LogUTCTime)
		ts = append(ts, t)
	}

	defer rows.Close()

	return ts, err
}

func getProgressesByUserIDSessionID(userid int, sessionid string, con *sql.DB) ([]int, error) {

	err := con.Ping()
	if err != nil {
		return nil, err
	}

	ps := make([]int,0)
	
	//fmt.Println(fmt.Sprintf("exec [dbo].[Select_ProgressByUserIDLessionID] ?,?"), userid, lessionid)
	rows, err := con.Query(fmt.Sprintf("exec [dbo].[Select_ProgressByUserIDSessionID] ?,?"), userid, sessionid)
	for rows.Next() {
		var p int
		err = rows.Scan(&p)
		ps = append(ps, p)
	}

	defer rows.Close()

	return ps, err
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


func getWordsBySemesterIDFromDB(t string, con *sql.DB) ([]Word, error) {

	err := con.Ping()
	if err != nil {
		return nil, err
	}

	var words []Word
	rows, err := con.Query("exec [dbo].[Select_WordsBySemesterID] '" + t + "'")
	for rows.Next() {
		var wd Word
		err = rows.Scan(&wd.ID, &wd.SemesterID, &wd.LessionID, &wd.ImagePath, &wd.AudioPath, &wd.Answer, &wd.Options)
		words = append(words, wd)
	}

	defer rows.Close()

	return words, err
}

func getOneWordByWordID(t int, con *sql.DB) ([]Word, error) {

	err := con.Ping()
	if err != nil {
		return nil, err
	}

	var words []Word
	rows, err := con.Query(fmt.Sprintf("exec [dbo].[Select_WordsByWordID] ?"), t)

	//rows, err := con.Query("exec [dbo].[Select_WordsByWordID] '" + t + "'")
	for rows.Next() {
		var wd Word
		err = rows.Scan(&wd.ID, &wd.SemesterID, &wd.LessionID, &wd.ImagePath, &wd.AudioPath, &wd.Answer, &wd.Options)
		words = append(words, wd)
	}

	defer rows.Close()

	return words, err
}


func getAnswerByUserIDSessionIDFromDB(userid int, sessionid string, con *sql.DB) ([]Answer, error) {

	err := con.Ping()
	if err != nil {
		return nil, err
	}

	var as []Answer
	
	rows, err := con.Query(fmt.Sprintf("exec [dbo].[Select_AnswerByUserIDSessionID] ?,?"), userid, sessionid)
	for rows.Next() {
		var a Answer
		err = rows.Scan(&a.AnswerID, &a.UserID, &a.SessionID, &a.SemesterID, &a.LessionID, &a.UserAnswer, &a.CorrectAnswer)
		as = append(as, a)
	}

	defer rows.Close()

	return as, err
}

func getComboFromDBBySessionID(sessionid string, con *sql.DB) ([]ComboRecord, error) {

	err := con.Ping()
	if err != nil {
		return nil, err
	}

	var cb []ComboRecord
	
	rows, err := con.Query(fmt.Sprintf("exec [dbo].[Select_ComboBySessionID] ?"), sessionid)
	for rows.Next() {
		var c ComboRecord
		err = rows.Scan(&c.UserName, &c.UserImgPath, &c.UserCombo)
		cb = append(cb, c)
	}

	defer rows.Close()

	return cb, err
}


func getPracticeByUserIDSessionIDFromDB(userid int, sessionid string, con *sql.DB) (Word, error) {

	err := con.Ping()
	var ww Word
	if err != nil {
		return ww, err
	}

	rows, err := con.Query(fmt.Sprintf("exec [dbo].[Select_PracticeWordByUserIDSessionID] ?,?"), userid, sessionid)

	for rows.Next() {
		var wd Word 	
		//fmt.Println(fmt.Sprintf("exec [dbo].[Select_PracticeWordByUserIDSessionID] ?,?"), userid, sessionid)

		err = rows.Scan(&wd.ID, &wd.SemesterID, &wd.LessionID, &wd.ImagePath, &wd.AudioPath, &wd.Answer, &wd.Options)
		ww = wd
	}

	defer rows.Close()

	return ww, err
}

func GetStairRecordsFromDBBySessionID(sessionid string, con *sql.DB) ([]StairRecord, error) {

	err := con.Ping()
	if err != nil {
		return nil, err
	}

	var sr []StairRecord
	
	rows, err := con.Query(fmt.Sprintf("exec [dbo].[Select_StairBySessionID] ?"), sessionid)
	for rows.Next() {
		var s StairRecord
		err = rows.Scan(&s.UserName, &s.LocalPhotoPath, &s.TotalCnt)
		sr = append(sr, s)
	}

	defer rows.Close()

	return sr, err
}


func getUserPracticeResultByUserIDSessionID(userid int, sessionid string, con *sql.DB) (PracticeResult, error) {

	err := con.Ping()
	var rt PracticeResult
	var comboCount = 0
	var tempComboCnt = 0
	var previousIsCorrect = true
	var cb = 0
 	if err != nil {
		return rt, err
	}
		

	rows, err := con.Query(fmt.Sprintf("exec [dbo].[Select_PracticeResultByUserIDSessionID] ?,?"), userid, sessionid)
	for rows.Next() {
		var pr PracticeRecord 	
//
		err = rows.Scan(&pr.UserPracticeID, &pr.UserID, &pr.SessionID, &pr.SemesterID, &pr.UserAnswer, &pr.CorrectAnswer, &pr.CorrectCnt, &pr.CorrectRate, pr.ComboInDB)
		rt.CorrectCnt = pr.CorrectCnt
		rt.CorrectRate = pr.CorrectRate
		cb = pr.ComboInDB
		if pr.UserAnswer == pr.CorrectAnswer {
			if previousIsCorrect {
				tempComboCnt += 1
			} else {
				tempComboCnt = 1
			}
			if tempComboCnt >= comboCount {
				comboCount = tempComboCnt
			}
			previousIsCorrect = true
		} else {
			tempComboCnt = 0
			previousIsCorrect = false
		}
			
	}

	
	
	rt.ComboCnt = comboCount
	if comboCount > cb {
			_, err = con.Exec(fmt.Sprintf("exec [dbo].[Upsert_ComboByUserIDComboCnt] ?,?,?"), userid,sessionid, comboCount)
	}
	defer rows.Close()
//fmt.Println(fmt.Sprintf("exec [dbo].[Select_PracticeResultByUserIDSessionID] ?,?"), userid, sessionid)

	return rt, err
}




func insertQuestion(q Question, con *sql.DB) (sql.Result, error) {

	err := con.Ping()
	if err != nil {
		return nil, err
	}

	r, err := con.Exec(fmt.Sprintf("exec [dbo].[Insert_Question] ?,?,?,?,?"), q.QuestionType, q.QuestionTitle, q.QuestionDescription, q.Answer, q.QuestionGroup)

	return r, err
}

func upsertAnswerToDB(a Answer,con *sql.DB)  (sql.Result, error) {

	err := con.Ping()
	if err != nil {
		return nil, err
	}
	//exec [dbo].[Upsert_Answer] '105554002','20180704',3,1,'qqq','ccc'
	//AsnwerID	int
    //UserID		int
    //SessionID	int
    //SemesterID	int
    //LessionID	int
    //UserAnswer	string
    //orrectAnswer string
	
	r, err := con.Exec(fmt.Sprintf("exec [dbo].[Upsert_Answer] ?,?,?,?,?,?"), a.UserID, a.SessionID, a.SemesterID, a.LessionID, a.UserAnswer, a.CorrectAnswer)

	return r, err
}



func insertOneUser(u User, con *sql.DB) (sql.Result, error) {

	err := con.Ping()
	if err != nil {
		return nil, err
	}

	r, err := con.Exec(fmt.Sprintf("exec [dbo].[Insert_User] ?,?,?,?"), u.UserID, u.UserName, u.PhotoPath, u.LocalPhotoPath)

	return r, err
}

func insertTracking(t Tracking, con *sql.DB) (sql.Result, error) {

	err := con.Ping()
	if err != nil {
		return nil, err
	}

	r, err := con.Exec(fmt.Sprintf("exec [dbo].[Insert_Tracking] ?,?,?,?,?,?"), t.XRequestID, t.SessionID, t.UserID, t.LessionID, t.Event, t.Note)

	return r, err
}

func insertProgress(t Progress, con *sql.DB) (sql.Result, error) {

	err := con.Ping()
	if err != nil {
		return nil, err
	}

	r, err := con.Exec(fmt.Sprintf("exec [dbo].[Insert_Progress] ?,?,?"), t.UserID, t.SessionID, t.LessionID)

	return r, err
}

func insertUserPractice(p PracticeRecord, con *sql.DB) (sql.Result, error) {

	err := con.Ping()
	if err != nil {
		return nil, err
	}

	r, err := con.Exec(fmt.Sprintf("exec [dbo].[Insert_PracticeRecord] ?,?,?,?,?"), p.UserID, p.SessionID, p.SemesterID, p.UserAnswer, p.CorrectAnswer)

	return r, err
}

func insertWord(wd Word, con *sql.DB) (sql.Result, error) {

	err := con.Ping()
	if err != nil {
		return nil, err
	}

	r, err := con.Exec(fmt.Sprintf("exec [dbo].[Insert_Word] ?,?,?,?,?,?"), wd.SemesterID, wd.LessionID, wd.ImagePath, wd.AudioPath, wd.Answer, wd.Options)

	return r, err
}


func getUserInfoByUserID(t string, con *sql.DB) (User, error) {

	var ts User
	err := con.Ping()
	if err != nil {
		return ts, err
	}

	rows, err := con.Query(fmt.Sprintf("exec [dbo].[Select_UserByUserID] ?"), t)
	for rows.Next() {
		var t User
		err = rows.Scan(&t.UserID, &t.UserName, &t.PhotoPath)
		ts = t
	}

	defer rows.Close()

	return ts, err
}
