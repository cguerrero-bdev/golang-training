package service

import (
	"fmt"

	"github.com/cguerrero-bdev/golang-training/final-project/api/components/definition/dao"
	"github.com/cguerrero-bdev/golang-training/final-project/api/components/definition/service"
	"github.com/cguerrero-bdev/golang-training/final-project/api/util"
)

type QuestionManager struct {
	QuestionDao dao.QuestionDao
	UserDao     dao.UserDao
}

func (questionManager *QuestionManager) GetQuestions() ([]service.Question, *util.ApplicationError) {

	questionEntities, err := questionManager.QuestionDao.GetQuestions()

	result := make([]service.Question, 0)
	for _, questionEntity := range questionEntities {
		result = append(result, createQuestion(&questionEntity, ""))

	}

	if err != nil {
		fmt.Printf("Error: %s -> %v\n", "UpdateQuestion", err)
	}

	return result, err
}

func (questionManager *QuestionManager) GetQuestionById(id int) (*service.Question, *error) {

	questionEntity, err := questionManager.QuestionDao.GetQuestionById(id)

	if err != nil {
		return nil, err
	}

	userEntity, err := questionManager.UserDao.GetUserById(questionEntity.Id)

	if err != nil {
		return nil, err
	}

	result := service.Question{
		Id:        questionEntity.Id,
		Statement: questionEntity.Statement,
		UserName:  userEntity.UserName,
	}

	if err != nil {
		fmt.Printf("Error: %s -> %v\n", "UpdateQuestion", err)
	}

	return &result, err
}

func (questionManager *QuestionManager) GetQuestionsByUserName(userName string) ([]service.Question, *error) {

	userEntity, err := questionManager.UserDao.GetUserByName(userName)

	if err != nil {

		return []service.Question{}, err
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

func (questionManager *QuestionManager) CreateQuestion(question *service.Question) (*service.Question, *error) {

	userEntity, err := questionManager.UserDao.GetUserByName(question.UserName)
	questionEntity := &dao.QuestionEntity{Id: question.Id, Statement: question.Statement, UserId: userEntity.Id}
	questionEntity, err = questionManager.QuestionDao.CreateQuestion(questionEntity)

	if err != nil {
		fmt.Printf("Error: %s -> %v\n", "CreateQuestion", err)
	}

	return question, err

}

func (questionManager *QuestionManager) UpdateQuestion(question *service.Question) (*service.Question, *util.ApplicationError) {

	questionEntity, err := questionManager.QuestionDao.GetQuestionById(question.Id)

	var applicationError *util.ApplicationError

	isThereAChange := question.Answer != questionEntity.Answer

	if isThereAChange {
		questionEntity.Answer = question.Answer

		answeredBy := question.AnsweredBy

		if answeredBy == "" {
			answeredBy = question.UserName
		}

		userEntity, err := questionManager.UserDao.GetUserByName(answeredBy)

		if err != nil {

			return question, util.GenerateApplicationError()
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

	if err != nil {
		util.GenerateApplicationError()
	}

	return question, applicationError
}

func (questionManager *QuestionManager) DeleteQuestion(id int) *error {

	err := questionManager.QuestionDao.DeleteQuestion(id)

	if err != nil {
		fmt.Printf("Error: %s -> %v\n", "UpdateQuestion", err)
	}

	return err
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
