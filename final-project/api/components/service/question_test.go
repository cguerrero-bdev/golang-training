package service

import (
	"testing"

	"github.com/cguerrero-bdev/golang-training/final-project/api/components/definition/dao"
	"github.com/cguerrero-bdev/golang-training/final-project/api/components/definition/service"
	mockdao "github.com/cguerrero-bdev/golang-training/final-project/api/components/mock/dao"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var userEntity = dao.UserEntity{
	Id:       2,
	UserName: "User2",
}

func TestGetQuestions(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	var questionEntities = []dao.QuestionEntity{
		{
			Id:        1,
			Statement: "Question 1",
			UserId:    1,
		},
		{
			Id:        2,
			Statement: "Question 2",
			UserId:    2,
		},
		{
			Id:        3,
			Statement: "Question 3",
			UserId:    1,
		},
	}

	mockQuestionDao := mockdao.NewMockQuestionDao(controller)
	mockQuestionDao.EXPECT().
		GetQuestions().
		Return(questionEntities, nil)

	mockUserDao := mockdao.NewMockUserDao(controller)

	questionService := QuestionManager{
		QuestionDao: mockQuestionDao,
		UserDao:     mockUserDao,
	}

	result, error := questionService.GetQuestions()

	assert.Nil(t, error, "Should not return an error")
	assert.Len(t, result, len(questionEntities))

	for i, questionEntity := range questionEntities {
		assert.Equal(t, result[i].Id, questionEntity.Id, "Id should be equal")
		assert.Equal(t, result[i].Statement, questionEntity.Statement, "Statement should be equal")
	}
}

func TestGetQuestionById(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	questionEntity := &dao.QuestionEntity{
		Id:        2,
		Statement: "Question 2",
		UserId:    2,
	}

	mockQuestionDao := mockdao.NewMockQuestionDao(controller)
	mockQuestionDao.EXPECT().
		GetQuestionById(2).
		Return(questionEntity, nil)

	mockUserDao := mockdao.NewMockUserDao(controller)
	mockUserDao.EXPECT().
		GetUserById(2).
		Return(&userEntity, nil)

	questionService := QuestionManager{
		QuestionDao: mockQuestionDao,
		UserDao:     mockUserDao,
	}

	result, error := questionService.GetQuestionById(2)

	assert.Nil(t, error, "Should not return an error")
	assert.Equal(t, result.Id, questionEntity.Id, "Id should be equal")
	assert.Equal(t, result.Statement, questionEntity.Statement, "Statement should be equal")
}

func TestService_CreateQuestion(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	questionEntity := dao.QuestionEntity{
		Id: 4, Statement: "Question 4", UserId: 2,
	}

	question := service.Question{
		Id: 4, Statement: "Question 4", UserName: "User2",
	}

	mockQuestionDao := mockdao.NewMockQuestionDao(controller)
	mockQuestionDao.EXPECT().
		CreateQuestion(&questionEntity).
		Return(&questionEntity, nil)

	mockUserDao := mockdao.NewMockUserDao(controller)
	mockUserDao.EXPECT().
		GetUserByName("User2").
		Return(&userEntity, nil)

	questionService := QuestionManager{
		QuestionDao: mockQuestionDao,
		UserDao:     mockUserDao,
	}

	result, error := questionService.CreateQuestion(&question)

	assert.Nil(t, error, "Should not return an error")
	assert.Equal(t, result.Id, questionEntity.Id, "Id should be equal")
	assert.Equal(t, result.Statement, questionEntity.Statement, "Statement should be equal")
}

func TestService_UpdateQuestion(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	questionEntity := dao.QuestionEntity{
		Id:         3,
		Statement:  "Question 3",
		UserId:     2,
		Answer:     "",
		AnsweredBy: 0,
	}

	question := service.Question{
		Id:         3,
		Statement:  "Question 3 updated",
		UserName:   "User2",
		Answer:     "Answer question 3",
		AnsweredBy: "User2",
	}

	userEntity = dao.UserEntity{
		Id:       2,
		UserName: "User2",
	}

	mockQuestionDao := mockdao.NewMockQuestionDao(controller)
	mockQuestionDao.EXPECT().
		GetQuestionById(3).
		Return(&questionEntity, nil)

	mockQuestionDao.EXPECT().
		UpdateQuestion(&questionEntity).
		Return(&questionEntity, nil)

	mockUserDao := mockdao.NewMockUserDao(controller)
	mockUserDao.EXPECT().
		GetUserByName("User2").
		Return(&userEntity, nil)

	questionService := QuestionManager{
		QuestionDao: mockQuestionDao,
		UserDao:     mockUserDao,
	}

	result, error := questionService.UpdateQuestion(&question)

	assert.Nil(t, error, "Should not return an error")
	assert.Equal(t, result.Id, question.Id, "Wrong value for 'Id' field")
	assert.Equal(t, result.Statement, question.Statement, "Wrong value for 'Statement' field")
	assert.Equal(t, result.Answer, question.Answer, "Wrong value for 'Answer' field")
}

func TestDeleteQuestion(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockQuestionDao := mockdao.NewMockQuestionDao(controller)
	mockQuestionDao.EXPECT().
		DeleteQuestion(1).
		Return(nil)

	mockUserDao := mockdao.NewMockUserDao(controller)

	questionService := QuestionManager{
		QuestionDao: mockQuestionDao,
		UserDao:     mockUserDao,
	}

	error := questionService.DeleteQuestion(1)
	assert.Nil(t, error, "Should not return an error")
}
