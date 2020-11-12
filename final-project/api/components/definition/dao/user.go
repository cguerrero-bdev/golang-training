package dao

type UserEntity struct {
	Id       int
	UserName string
}

type UserDao interface {
	GetUserByName(userName string) (*UserEntity, *error)
	GetUserById(id int) (*UserEntity, *error)
}
