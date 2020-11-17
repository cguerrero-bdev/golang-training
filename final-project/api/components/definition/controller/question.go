package controller

import "net/http"

type JsonQuestion struct {
	Id         string `json:"id"`
	Statement  string `json:"statement"`
	UserName   string `json:"username"`
	Answer     string `json:"answer"`
	AnsweredBy string `json:"answeredby"`
}

type QuestionController interface {
	GetQuestions(w http.ResponseWriter, r *http.Request)
	GetQuestionById(w http.ResponseWriter, r *http.Request)
	GetQuestionsByUserName(w http.ResponseWriter, r *http.Request)
	CreateQuestion(w http.ResponseWriter, r *http.Request)
	UpdateQuestion(w http.ResponseWriter, r *http.Request)
	DeleteQuestion(w http.ResponseWriter, r *http.Request)
}
