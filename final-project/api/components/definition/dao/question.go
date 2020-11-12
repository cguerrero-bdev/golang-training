package dao

import (
	"github.com/cguerrero-bdev/golang-training/final-project/api/util"
)

type QuestionEntity struct {
	Id         int
	Statement  string
	UserId     int
	Answer     string
	AnsweredBy int
}

type QuestionDao interface {
	GetQuestions() ([]QuestionEntity, *util.ApplicationError)
	GetQuestionById(id int) (*QuestionEntity, *error)
	CreateQuestion(q *QuestionEntity) (*QuestionEntity, *error)
	GetQuestionsByUserId(id int) ([]QuestionEntity, *error)
	UpdateQuestion(q *QuestionEntity) (*QuestionEntity, *util.ApplicationError)
	DeleteQuestion(id int) *error
}
