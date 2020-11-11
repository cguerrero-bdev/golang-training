package service

import "github.com/cguerrero-bdev/golang-training/final-project/api/util"

type Question struct {
	Id         int
	Statement  string
	UserName   string
	Answere    string
	AnsweredBy string
}

type QuestionManager interface {
	GetQuestions() ([]Question, *util.ApplicationError)
	GetQuestionById(id int) (Question, error)
	GetQuestionsByUserName(userName string) ([]Question, error)
	CreateQuestion(question Question) (Question, error)
	UpdateQuestion(question Question) (Question, *util.ApplicationError)
	DeleteQuestion(id int) error
}
