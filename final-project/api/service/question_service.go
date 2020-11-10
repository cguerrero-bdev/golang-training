package service

import (
	"fmt"

	"github.com/cguerrero-bdev/golang-training/final-project/api/dao"
	"github.com/cguerrero-bdev/golang-training/final-project/api/util"
)

type Question struct {
	Id         int
	Statement  string
	UserName   string
	Answere    string
	AnsweredBy string
}

type QuestionManager struct {
	QuestionRepository dao.QuestionDao
	UserRepository     dao.UserDao
}

func (questionManager *QuestionManager) GetQuestions() ([]Question, *util.ApplicationError) {

	questionEntities, err := questionManager.QuestionRepository.GetQuestions()

	result := make([]Question, 0)
	for _, questionEntity := range questionEntities {
		result = append(result, createQuestion(&questionEntity, ""))

	}

	if err != nil {
		fmt.Printf("Error: %s -> %v\n", "UpdateQuestion", err)
	}

	return result, err
}

func (questionManager *QuestionManager) GetQuestionById(id int) (Question, error) {

	questionEntity, err := questionManager.QuestionRepository.GetQuestionById(id)

	if err != nil {
		return Question{}, err
	}

	userEntity, err := questionManager.UserRepository.GetUserById(questionEntity.Id)

	if err != nil {
		return Question{}, err
	}

	result := Question{
		Id:        questionEntity.Id,
		Statement: questionEntity.Statement,
		UserName:  userEntity.UserName,
	}

	if err != nil {
		fmt.Printf("Error: %s -> %v\n", "UpdateQuestion", err)
	}

	return result, err
}

func (questionManager *QuestionManager) GetQuestionsByUserName(userName string) ([]Question, error) {

	userEntity, err := questionManager.UserRepository.GetUserByName(userName)

	if err != nil {

		return []Question{}, err
	}

	questionEntities, err := questionManager.QuestionRepository.GetQuestionsByUserId(userEntity.Id)

	if err != nil {
		return []Question{}, err
	}

	result := make([]Question, 0)
	for _, questionEntity := range questionEntities {
		result = append(result, createQuestion(&questionEntity, userEntity.UserName))

	}

	if err != nil {
		fmt.Printf("Error: %s -> %v\n", "UpdateQuestion", err)
	}

	return result, err
}

func (questionManager *QuestionManager) CreateQuestion(question Question) (Question, error) {

	userEntity, err := questionManager.UserRepository.GetUserByName(question.UserName)
	questionEntity := dao.QuestionEntity{Id: question.Id, Statement: question.Statement, UserId: userEntity.Id}
	questionEntity, err = questionManager.QuestionRepository.CreateQuestion(questionEntity)

	if err != nil {
		fmt.Printf("Error: %s -> %v\n", "CreateQuestion", err)
	}

	return question, err

}

func (questionManager *QuestionManager) UpdateQuestion(question Question) (Question, *util.ApplicationError) {

	questionEntity, err := questionManager.QuestionRepository.GetQuestionById(question.Id)

	var applicationError *util.ApplicationError

	isThereAChange := question.Answere != questionEntity.Answere

	if isThereAChange {
		questionEntity.Answere = question.Answere

		answeredBy := question.AnsweredBy

		if answeredBy == "" {
			answeredBy = question.UserName
		}

		userEntity, err := questionManager.UserRepository.GetUserByName(answeredBy)

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

		questionEntity, applicationError = questionManager.QuestionRepository.UpdateQuestion(questionEntity)
	}

	if err != nil {
		util.GenerateApplicationError()
	}

	return question, applicationError
}

func (questionManager *QuestionManager) DeleteQuestion(id int) error {

	err := questionManager.QuestionRepository.DeleteQuestion(id)

	if err != nil {
		fmt.Printf("Error: %s -> %v\n", "UpdateQuestion", err)
	}

	return err
}

func createQuestion(questionEntity *dao.QuestionEntity, userName string) Question {

	result := Question{
		Id:         questionEntity.Id,
		Statement:  questionEntity.Statement,
		UserName:   userName,
		Answere:    questionEntity.Answere,
		AnsweredBy: userName,
	}

	return result
}
