package service

import "github.com/cguerrero-bdev/golang-training/final-project/api/util"

type Question struct {
	Id         int
	Statement  string
	UserName   string
	Answer     string
	AnsweredBy string
}

type QuestionManager interface {
	GetQuestions() ([]Question, util.ApplicationError)
	GetQuestionById(id int) (*Question, util.ApplicationError)
	GetQuestionsByUserName(userName string) ([]Question, util.ApplicationError)
	CreateQuestion(question *Question) (*Question, util.ApplicationError)
	UpdateQuestion(question *Question) (*Question, util.ApplicationError)
	DeleteQuestion(id int) util.ApplicationError
}
