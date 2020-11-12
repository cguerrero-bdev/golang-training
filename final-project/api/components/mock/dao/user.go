// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cguerrero-bdev/golang-training/final-project/api/components/definition/dao (interfaces: UserDao)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	dao "github.com/cguerrero-bdev/golang-training/final-project/api/components/definition/dao"
	gomock "github.com/golang/mock/gomock"
)

// MockUserDao is a mock of UserDao interface.
type MockUserDao struct {
	ctrl     *gomock.Controller
	recorder *MockUserDaoMockRecorder
}

// MockUserDaoMockRecorder is the mock recorder for MockUserDao.
type MockUserDaoMockRecorder struct {
	mock *MockUserDao
}

// NewMockUserDao creates a new mock instance.
func NewMockUserDao(ctrl *gomock.Controller) *MockUserDao {
	mock := &MockUserDao{ctrl: ctrl}
	mock.recorder = &MockUserDaoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserDao) EXPECT() *MockUserDaoMockRecorder {
	return m.recorder
}

// GetUserById mocks base method.
func (m *MockUserDao) GetUserById(arg0 int) (dao.UserEntity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserById", arg0)
	ret0, _ := ret[0].(dao.UserEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserById indicates an expected call of GetUserById.
func (mr *MockUserDaoMockRecorder) GetUserById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserById", reflect.TypeOf((*MockUserDao)(nil).GetUserById), arg0)
}

// GetUserByName mocks base method.
func (m *MockUserDao) GetUserByName(arg0 string) (dao.UserEntity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByName", arg0)
	ret0, _ := ret[0].(dao.UserEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByName indicates an expected call of GetUserByName.
func (mr *MockUserDaoMockRecorder) GetUserByName(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByName", reflect.TypeOf((*MockUserDao)(nil).GetUserByName), arg0)
}
