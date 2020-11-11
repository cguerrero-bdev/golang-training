package dao

import (
	"context"

	"github.com/cguerrero-bdev/golang-training/final-project/api/components/definition/dao"
	"github.com/jackc/pgx/v4"
)

type UserDao struct {
	Connection *pgx.Conn
}

const userSelect = "select id, user_name from app_user "

func (userDao *UserDao) GetUserByName(userName string) (dao.UserEntity, error) {

	userEntity := dao.UserEntity{}
	err := userDao.Connection.QueryRow(context.Background(), userSelect+" where user_name=$1", userName).Scan(&userEntity.Id, &userEntity.UserName)

	return userEntity, err
}

func (userDao *UserDao) GetUserById(id int) (dao.UserEntity, error) {

	userEntity := dao.UserEntity{}
	err := userDao.Connection.QueryRow(context.Background(), userSelect+" where id=$1", id).Scan(&userEntity.Id, &userEntity.UserName)

	return userEntity, err
}
