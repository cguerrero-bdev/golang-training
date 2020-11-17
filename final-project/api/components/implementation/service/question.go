package service

import (
	"fmt"
	"log"

	"github.com/cguerrero-bdev/golang-training/final-project/api/components/definition/dao"
	"github.com/cguerrero-bdev/golang-training/final-project/api/components/definition/service"
)

type QuestionManager struct {
	QuestionDao dao.QuestionDao
	UserDao     dao.UserDao
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
}

func (questionManager *QuestionManager) GetQuestions() ([]service.Question, error) {

	questionEntities, error := questionManager.QuestionDao.GetQuestions()

	if error != nil {
		return nil, error
	}

	result := make([]service.Question, 0)
	for _, questionEntity := range questionEntities {
		result = append(result, createQuestion(&questionEntity, ""))

	}

	return result, nil
}

func (questionManager *QuestionManager) GetQuestionById(id int) (*service.Question, error) {

	questionEntity, error := questionManager.QuestionDao.GetQuestionById(id)

	if error != nil {
		return nil, error
	}

	userEntity, error := questionManager.UserDao.GetUserById(questionEntity.UserId)

	if error != nil {
		return nil, error
	}

	result := service.Question{
		Id:        questionEntity.Id,
		Statement: questionEntity.Statement,
		UserName:  userEntity.UserName,
	}

	return &result, nil
}

func (questionManager *QuestionManager) GetQuestionsByUserName(userName string) ([]service.Question, error) {

	userEntity, error := questionManager.UserDao.GetUserByName(userName)

	if error != nil {

		return nil, error
	}

	questionEntities, err := questionManager.QuestionDao.GetQuestionsByUserId(userEntity.Id)

	if err != nil {
		return []service.Question{}, err
	}

	result := make([]service.Question, 0)
	for _, questionEntity := range questionEntities {
		result = append(result, createQuestion(&questionEntity, userEntity.UserName))

	}

	if err != nil {
		fmt.Printf("Error: %s -> %v\n", "UpdateQuestion", err)
	}

	return result, err
}

func (questionManager *QuestionManager) CreateQuestion(question *service.Question) (*service.Question, error) {

	userEntity, error := questionManager.UserDao.GetUserByName(question.UserName)

	if error != nil {
		return nil, error
	}

	questionEntity := &dao.QuestionEntity{Id: question.Id, Statement: question.Statement, UserId: userEntity.Id}
	questionEntity, error = questionManager.QuestionDao.CreateQuestion(questionEntity)

	if error != nil {
		return nil, error
	}

	return question, nil

}

func (questionManager *QuestionManager) UpdateQuestion(question *service.Question) (*service.Question, error) {

	questionEntity, error := questionManager.QuestionDao.GetQuestionById(question.Id)

	if error != nil {

		return question, error
	}

	isThereAChange := question.Answer != questionEntity.Answer

	if isThereAChange {
		questionEntity.Answer = question.Answer

		answeredBy := question.AnsweredBy

		if answeredBy == "" {
			answeredBy = question.UserName
		}

		userEntity, error := questionManager.UserDao.GetUserByName(answeredBy)

		if error != nil {

			return question, error
		}

		questionEntity.AnsweredBy = userEntity.Id
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

func (questionManager *QuestionManager) DeleteQuestion(id int) error {

	error := questionManager.QuestionDao.DeleteQuestion(id)

	if error != nil {
		return error
	}

	return nil
}

func createQuestion(questionEntity *dao.QuestionEntity, userName string) service.Question {

	result := service.Question{
		Id:         questionEntity.Id,
		Statement:  questionEntity.Statement,
		UserName:   userName,
		Answer:     questionEntity.Answer,
		AnsweredBy: userName,
	}

	return result
}
