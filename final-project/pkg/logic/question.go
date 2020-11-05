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

func (questionManager *QuestionManager) GetQuestions() {

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

func (questionManager *QuestionManager) GetQuestionByUserId(id int) {

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
