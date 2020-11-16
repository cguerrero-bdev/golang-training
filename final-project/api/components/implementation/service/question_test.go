package service

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/cguerrero-bdev/golang-training/final-project/api/components/definition/dao"
	"github.com/cguerrero-bdev/golang-training/final-project/api/components/definition/service"

	mockdao "github.com/cguerrero-bdev/golang-training/final-project/api/components/mock/dao"
)

var userEntity = dao.UserEntity{
	Id:       2,
	UserName: "User2",
}

func TestGetQuestions(t *testing.T) {

	var questionEntities = []dao.QuestionEntity{
		{
			Id: 1, Statement: "Question 1", UserId: 1,
		},
		{
			Id: 2, Statement: "Question 2", UserId: 2,
		},
		{
			Id: 3, Statement: "Question 3", UserId: 1,
		},
	}

	controller := gomock.NewController(t)
	defer controller.Finish()

	mockQuestionDao := mockdao.NewMockQuestionDao(controller)
	mockUserDao := mockdao.NewMockUserDao(controller)

	mockQuestionDao.EXPECT().GetQuestions().Return(questionEntities, nil)

	questionService := QuestionManager{
		QuestionDao: mockQuestionDao,
		UserDao:     mockUserDao,
	}

	result, error := questionService.GetQuestions()

	assert := assert.New(t)
	assert.Nil(error, "Should not return an error")
	assert.Len(result, len(questionEntities))

	for i, questionEntity := range questionEntities {
		assert.Equal(result[i].Id, questionEntity.Id, "Id should be equal")
		assert.Equal(result[i].Statement, questionEntity.Statement, "Statement should be equal")
	}
}

func TestGetQuestionById(t *testing.T) {

	questionEntities := dao.QuestionEntity{
		Id: 2, Statement: "Question 2", UserId: 2,
	}

	controller := gomock.NewController(t)
	defer controller.Finish()

	mockQuestionDao := mockdao.NewMockQuestionDao(controller)
	mockUserDao := mockdao.NewMockUserDao(controller)

	mockQuestionDao.EXPECT().GetQuestionById(2).Return(&questionEntities, nil)
	mockUserDao.EXPECT().GetUserById(2).Return(userEntity, nil)

	questionService := QuestionManager{
		QuestionDao: mockQuestionDao,
		UserDao:     mockUserDao,
	}

	result, error := questionService.GetQuestionById(2)

	assert := assert.New(t)
	assert.Nil(error, "Should not return an error")
	assert.Equal(result.Id, questionEntities.Id, "Id should be equal")
	assert.Equal(result.Statement, questionEntities.Statement, "Statement should be equal")

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
	mockUserDao := mockdao.NewMockUserDao(controller)

	mockQuestionDao.EXPECT().
		CreateQuestion(&questionEntity).
		Return(&questionEntity, nil)

	mockUserDao.EXPECT().GetUserByName("User2").Return(&userEntity, nil)

	questionService := QuestionManager{
		QuestionDao: mockQuestionDao,
		UserDao:     mockUserDao,
	}

	result, error := questionService.CreateQuestion(&question)

	assert := assert.New(t)
	assert.Nil(error, "Should not return an error")
	assert.Equal(result.Id, questionEntity.Id, "Id should be equal")
	assert.Equal(result.Statement, questionEntity.Statement, "Statement should be equal")

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
	mockUserDao := mockdao.NewMockUserDao(controller)

	mockUserDao.EXPECT().GetUserByName("User2").Return(&userEntity, nil)

	mockQuestionDao.EXPECT().GetQuestionById(3).Return(&questionEntity, nil)

	mockQuestionDao.EXPECT().
		UpdateQuestion(&questionEntity).
		Return(&questionEntity, nil)

	questionService := QuestionManager{
		QuestionDao: mockQuestionDao,
		UserDao:     mockUserDao,
	}

	result, error := questionService.UpdateQuestion(&question)

	assert := assert.New(t)
	assert.Nil(error, "Should not return an error")
	assert.Equal(result.Id, question.Id, "Wrong value for 'Id' field")
	assert.Equal(result.Statement, question.Statement, "Wrong value for 'Statement' field")
	assert.Equal(result.Answer, question.Answer, "Wrong value for 'Answer' field")

}

func TestDeleteQuestion(t *testing.T) {

	controller := gomock.NewController(t)
	defer controller.Finish()

	mockQuestionDao := mockdao.NewMockQuestionDao(controller)
	mockUserDao := mockdao.NewMockUserDao(controller)

	mockQuestionDao.EXPECT().DeleteQuestion(1).Return(nil)

	questionService := QuestionManager{
		QuestionDao: mockQuestionDao,
		UserDao:     mockUserDao,
	}

	error := questionService.DeleteQuestion(1)

	assert := assert.New(t)
	assert.Nil(error, "Should not return an error")

}
