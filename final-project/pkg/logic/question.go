package logic

import (
	"github.com/cguerrero-bdev/golang-training/final-project/pkg/persistence"
)

type Question struct {
	Id       int
	Text     string
	UserName string
}

type QuestionManager struct {
	QuestionRepository persistence.QuestionRepository
	UserRepository     persistence.UserRepository
}

func (questionManager *QuestionManager) GetQuestions() ([]Question, error) {

	questionEntities, err := questionManager.QuestionRepository.GetQuestions()

	if err != nil {
		return []Question{}, err
	}

	result := make([]Question, 0)
	for _, questionEntity := range questionEntities {
		result = append(result, createQuestion(&questionEntity, ""))

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
		Id:       questionEntity.Id,
		Text:     questionEntity.Text,
		UserName: userEntity.UserName,
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

	return result, err
}

func (questionManager *QuestionManager) CreateQuestion(question Question) (Question, error) {

	userEntity, err := questionManager.UserRepository.GetUserByName(question.UserName)

	questionEntity := persistence.QuestionEntity{Id: question.Id, Text: question.Text, UserId: userEntity.Id}
	questionEntity, err = questionManager.QuestionRepository.CreateQuestion(questionEntity)

	return question, err

}

func (questionManager *QuestionManager) UpdateQuestion(question Question) {

}

func (questionManager *QuestionManager) DeleteQuestion(id int) {

}

func createQuestion(questionEntity *persistence.QuestionEntity, userName string) Question {

	result := Question{
		Id:       questionEntity.Id,
		Text:     questionEntity.Text,
		UserName: userName,
	}

	return result
}
