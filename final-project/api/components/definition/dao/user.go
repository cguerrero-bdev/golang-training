package dao

import "github.com/cguerrero-bdev/golang-training/final-project/api/util"

type UserEntity struct {
	Id       int
	UserName string
}

type UserDao interface {
	GetUserByName(userName string) (*UserEntity, util.ApplicationError)
	GetUserById(id int) (*UserEntity, util.ApplicationError)
}
