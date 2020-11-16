package dao

import (
	"context"
	"log"

	"github.com/cguerrero-bdev/golang-training/final-project/api/components/definition/dao"
	"github.com/jackc/pgx/v4"
)

type UserDao struct {
	Connection *pgx.Conn

	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
}

const userSelect = "select id, user_name from app_user "

func (userDao *UserDao) GetUserByName(userName string) (*dao.UserEntity, error) {

	userEntity := dao.UserEntity{}
	err := userDao.Connection.QueryRow(context.Background(), userSelect+" where user_name=$1", userName).Scan(&userEntity.Id, &userEntity.UserName)
	if err != nil {
		userDao.ErrorLogger.Println(err.Error())
		return nil, err
	}

	return &userEntity, nil
}

func (userDao *UserDao) GetUserById(id int) (*dao.UserEntity, error) {

	userEntity := dao.UserEntity{}
	err := userDao.Connection.QueryRow(context.Background(), userSelect+" where id=$1", id).Scan(&userEntity.Id, &userEntity.UserName)
	if err != nil {
		userDao.ErrorLogger.Println(err.Error())
		return nil, err
	}

	return &userEntity, nil
}
