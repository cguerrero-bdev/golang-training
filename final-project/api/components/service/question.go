package service

import (
	"log"

	"github.com/cguerrero-bdev/golang-training/final-project/api/components/dao"
	"github.com/cguerrero-bdev/golang-training/final-project/api/components/model"
)

type QuestionManager interface {
	GetQuestions() ([]model.Question, error)
	GetQuestionById(id int) (*model.Question, error)
	GetQuestionsByUserId(id int) ([]model.Question, error)
	CreateQuestion(question *model.Question) (*model.Question, error)
	UpdateQuestion(question *model.Question) (*model.Question, error)
	DeleteQuestion(id int) error
}

type QuestionManagerImplementation struct {
	QuestionDao dao.QuestionDao
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
}

func (questionManager *QuestionManagerImplementation) GetQuestions() ([]model.Question, error) {

	result, error := questionManager.QuestionDao.GetQuestions()

	if error != nil {
		return nil, error
	}

	return result, nil
}

func (questionManager *QuestionManagerImplementation) GetQuestionById(id int) (*model.Question, error) {

	result, error := questionManager.QuestionDao.GetQuestionById(id)

	if error != nil {
		return nil, error
	}

	return result, nil
}

func (questionManager *QuestionManagerImplementation) GetQuestionsByUserId(id int) ([]model.Question, error) {

	result, err := questionManager.QuestionDao.GetQuestionsByUserId(id)

	if err != nil {
		return []model.Question{}, err
	}

	return result, err
}

func (questionManager *QuestionManagerImplementation) CreateQuestion(question *model.Question) (*model.Question, error) {

	question, error := questionManager.QuestionDao.CreateQuestion(question)

	if error != nil {
		return nil, error
	}

	return question, nil

}

func (questionManager *QuestionManagerImplementation) UpdateQuestion(question *model.Question) (*model.Question, error) {

	questionEntity, error := questionManager.QuestionDao.GetQuestionById(question.Id)

	if error != nil {

		return question, error
	}

	isThereAChange := question.Answer != questionEntity.Answer

	if isThereAChange {
		questionEntity.Answer = question.Answer

		if question.AnsweredBy == 0 {
			question.AnsweredBy = question.UserId
		}

	}

	if question.Statement != questionEntity.Statement {

		questionEntity.Statement = question.Statement
		isThereAChange = true

	}

	if isThereAChange {

		questionEntity, error = questionManager.QuestionDao.UpdateQuestion(questionEntity)
	}

	if error != nil {
		return nil, error
	}

	return question, error
}

func (questionManager *QuestionManagerImplementation) DeleteQuestion(id int) error {

	error := questionManager.QuestionDao.DeleteQuestion(id)

	if error != nil {
		return error
	}

	return nil
}
