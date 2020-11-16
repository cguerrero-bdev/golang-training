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

	questionEntities, applicationError := questionManager.QuestionDao.GetQuestions()

	if applicationError != nil {
		return nil, applicationError
	}

	result := make([]service.Question, 0)
	for _, questionEntity := range questionEntities {
		result = append(result, createQuestion(&questionEntity, ""))

	}

	return result, nil
}

func (questionManager *QuestionManager) GetQuestionById(id int) (*service.Question, error) {

	questionEntity, applicationError := questionManager.QuestionDao.GetQuestionById(id)

	if applicationError != nil {
		return nil, applicationError
	}

	userEntity, applicationError := questionManager.UserDao.GetUserById(questionEntity.Id)

	if applicationError != nil {
		return nil, applicationError
	}

	result := service.Question{
		Id:        questionEntity.Id,
		Statement: questionEntity.Statement,
		UserName:  userEntity.UserName,
	}

	return &result, nil
}

func (questionManager *QuestionManager) GetQuestionsByUserName(userName string) ([]service.Question, error) {

	userEntity, applicationError := questionManager.UserDao.GetUserByName(userName)

	if applicationError != nil {

		return nil, applicationError
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

	userEntity, applicationError := questionManager.UserDao.GetUserByName(question.UserName)

	if applicationError != nil {
		return nil, applicationError
	}

	questionEntity := &dao.QuestionEntity{Id: question.Id, Statement: question.Statement, UserId: userEntity.Id}
	questionEntity, applicationError = questionManager.QuestionDao.CreateQuestion(questionEntity)

	if applicationError != nil {
		return nil, applicationError
	}

	return question, nil

}

func (questionManager *QuestionManager) UpdateQuestion(question *service.Question) (*service.Question, error) {

	questionEntity, applicationError := questionManager.QuestionDao.GetQuestionById(question.Id)

	if applicationError != nil {

		return question, applicationError
	}

	isThereAChange := question.Answer != questionEntity.Answer

	if isThereAChange {
		questionEntity.Answer = question.Answer

		answeredBy := question.AnsweredBy

		if answeredBy == "" {
			answeredBy = question.UserName
		}

		userEntity, applicationError := questionManager.UserDao.GetUserByName(answeredBy)

		if applicationError != nil {

			return question, applicationError
		}

		questionEntity.AnsweredBy = userEntity.Id
	}

	if question.Statement != questionEntity.Statement {

		questionEntity.Statement = question.Statement
		isThereAChange = true

	}

	if isThereAChange {

		questionEntity, applicationError = questionManager.QuestionDao.UpdateQuestion(questionEntity)
	}

	if applicationError != nil {
		return nil, applicationError
	}

	return question, applicationError
}

func (questionManager *QuestionManager) DeleteQuestion(id int) error {

	applicationError := questionManager.QuestionDao.DeleteQuestion(id)

	if applicationError != nil {
		return applicationError
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
