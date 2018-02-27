# POC

The project will demonstrate a simple web api 

## Prerequisites

Install 3rd party library

```
go get -u github.com/denisenkom/go-mssqldb

go get -u github.com/gorilla/mux

go get -u github.com/spf13/viper
```

Usage

```
Request sample
GET/
	http://xx.xx.xx.xx:8000/question/{groupid}
	http://xx.xx.xx.xx:8000/tracking/{userid}
	router.HandleFunc("/question", CreateQuestion).Methods("POST")
	router.HandleFunc("/tracking", CreateTracking).Methods("POST")
POST
	http://xx.xx.xx.xx:8000/question
	body:
		{"ID":2,"QuestionType":"1","QuestionTitle":"What's the mid day of a week?","QuestionDescription":null,"Answer":"Saturday","QuestionGroup":1}
		
	http://xx.xx.xx.xx:8000/tracking
		{"TrackingID":2,"UserID":1,"SessionID":"8E19FE65-7E74-4338-BFDD-556D0C5D709F","Stage":"1","Msg":"Q1_Start","LogTime":"2018-02-01T09:40:04.233Z","QuestionID":1,"Answer":"test answer"}
	
```
