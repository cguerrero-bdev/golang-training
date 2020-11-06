package persistence

import (
	"context"

	"github.com/jackc/pgx/v4"
)

type UserEntity struct {
	Id       int
	UserName string
}

type UserRepository struct {
	Connection *pgx.Conn
}

const userSelect = "select id, user_name from app_user "

func (userRepository *UserRepository) GetUserByName(userName string) (UserEntity, error) {

	userEntity := UserEntity{}
	err := userRepository.Connection.QueryRow(context.Background(), userSelect+" where user_name=$1", userName).Scan(&userEntity.Id, &userEntity.UserName)

	return userEntity, err
}

func (userRepository *UserRepository) GetUserById(id int) (UserEntity, error) {

	userEntity := UserEntity{}
	err := userRepository.Connection.QueryRow(context.Background(), userSelect+" where id=$1", id).Scan(&userEntity.Id, &userEntity.UserName)

	return userEntity, err
}
