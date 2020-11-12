package service

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/cguerrero-bdev/golang-training/final-project/api/components/definition/dao"
	"github.com/cguerrero-bdev/golang-training/final-project/api/components/definition/service"

	mockdao "github.com/cguerrero-bdev/golang-training/final-project/api/components/mock/dao"
)

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

func TestGetQuestions(t *testing.T) {

	controller := gomock.NewController(t)
	defer controller.Finish()

	mockQuestionDao := mockdao.NewMockQuestionDao(controller)

	mockQuestionDao.EXPECT().GetQuestions().Return(questionEntities, nil)

	questionService := QuestionManager{
		QuestionDao: mockQuestionDao,
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

	controller := gomock.NewController(t)
	defer controller.Finish()

	userEntity := dao.UserEntity{
		Id: 2,
	}

	mockQuestionDao := mockdao.NewMockQuestionDao(controller)
	mockUserDao := mockdao.NewMockUserDao(controller)

	mockQuestionDao.EXPECT().GetQuestionById(2).Return(&questionEntities[1], nil)
	mockUserDao.EXPECT().GetUserById(2).Return(userEntity, nil)

	questionService := QuestionManager{
		QuestionDao: mockQuestionDao,
		UserDao:     mockUserDao,
	}

	result, error := questionService.GetQuestionById(2)

	assert := assert.New(t)
	assert.Nil(error, "Should not return an error")
	assert.Equal(result.Id, questionEntities[1].Id, "Id should be equal")
	assert.Equal(result.Statement, questionEntities[1].Statement, "Statement should be equal")

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

	userEntity := dao.UserEntity{
		Id:       2,
		UserName: "User2",
	}

	mockQuestionDao := mockdao.NewMockQuestionDao(controller)
	mockUserDao := mockdao.NewMockUserDao(controller)

	mockQuestionDao.EXPECT().
		CreateQuestion(&questionEntity).
		Return(&questionEntity, nil)

	mockUserDao.EXPECT().GetUserByName("User2").Return(userEntity, nil)

	questionService := QuestionManager{
		QuestionDao: mockQuestionDao,
		UserDao:     mockUserDao,
	}

	result, error := questionService.CreateQuestion(question)

	assert := assert.New(t)
	assert.Nil(error, "Should not return an error")
	assert.Equal(result.Id, questionEntity.Id, "Id should be equal")
	assert.Equal(result.Statement, questionEntity.Statement, "Statement should be equal")

}