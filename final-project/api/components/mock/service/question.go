// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cguerrero-bdev/golang-training/final-project/api/components/definition/service (interfaces: QuestionManager)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	service "github.com/cguerrero-bdev/golang-training/final-project/api/components/definition/service"
	util "github.com/cguerrero-bdev/golang-training/final-project/api/util"
	gomock "github.com/golang/mock/gomock"
)

// MockQuestionManager is a mock of QuestionManager interface.
type MockQuestionManager struct {
	ctrl     *gomock.Controller
	recorder *MockQuestionManagerMockRecorder
}

// MockQuestionManagerMockRecorder is the mock recorder for MockQuestionManager.
type MockQuestionManagerMockRecorder struct {
	mock *MockQuestionManager
}

// NewMockQuestionManager creates a new mock instance.
func NewMockQuestionManager(ctrl *gomock.Controller) *MockQuestionManager {
	mock := &MockQuestionManager{ctrl: ctrl}
	mock.recorder = &MockQuestionManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQuestionManager) EXPECT() *MockQuestionManagerMockRecorder {
	return m.recorder
}

// CreateQuestion mocks base method.
func (m *MockQuestionManager) CreateQuestion(arg0 service.Question) (service.Question, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateQuestion", arg0)
	ret0, _ := ret[0].(service.Question)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateQuestion indicates an expected call of CreateQuestion.
func (mr *MockQuestionManagerMockRecorder) CreateQuestion(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateQuestion", reflect.TypeOf((*MockQuestionManager)(nil).CreateQuestion), arg0)
}

// DeleteQuestion mocks base method.
func (m *MockQuestionManager) DeleteQuestion(arg0 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteQuestion", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteQuestion indicates an expected call of DeleteQuestion.
func (mr *MockQuestionManagerMockRecorder) DeleteQuestion(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteQuestion", reflect.TypeOf((*MockQuestionManager)(nil).DeleteQuestion), arg0)
}

// GetQuestionById mocks base method.
func (m *MockQuestionManager) GetQuestionById(arg0 int) (service.Question, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetQuestionById", arg0)
	ret0, _ := ret[0].(service.Question)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetQuestionById indicates an expected call of GetQuestionById.
func (mr *MockQuestionManagerMockRecorder) GetQuestionById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQuestionById", reflect.TypeOf((*MockQuestionManager)(nil).GetQuestionById), arg0)
}

// GetQuestions mocks base method.
func (m *MockQuestionManager) GetQuestions() ([]service.Question, *util.ApplicationError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetQuestions")
	ret0, _ := ret[0].([]service.Question)
	ret1, _ := ret[1].(*util.ApplicationError)
	return ret0, ret1
}

// GetQuestions indicates an expected call of GetQuestions.
func (mr *MockQuestionManagerMockRecorder) GetQuestions() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQuestions", reflect.TypeOf((*MockQuestionManager)(nil).GetQuestions))
}

// GetQuestionsByUserName mocks base method.
func (m *MockQuestionManager) GetQuestionsByUserName(arg0 string) ([]service.Question, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetQuestionsByUserName", arg0)
	ret0, _ := ret[0].([]service.Question)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetQuestionsByUserName indicates an expected call of GetQuestionsByUserName.
func (mr *MockQuestionManagerMockRecorder) GetQuestionsByUserName(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQuestionsByUserName", reflect.TypeOf((*MockQuestionManager)(nil).GetQuestionsByUserName), arg0)
}

// UpdateQuestion mocks base method.
func (m *MockQuestionManager) UpdateQuestion(arg0 service.Question) (service.Question, *util.ApplicationError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateQuestion", arg0)
	ret0, _ := ret[0].(service.Question)
	ret1, _ := ret[1].(*util.ApplicationError)
	return ret0, ret1
}

// UpdateQuestion indicates an expected call of UpdateQuestion.
func (mr *MockQuestionManagerMockRecorder) UpdateQuestion(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateQuestion", reflect.TypeOf((*MockQuestionManager)(nil).UpdateQuestion), arg0)
}