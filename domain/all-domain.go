package domain

import "github.com/arifwidiasan/todo-app/model"

type AdapterRepository interface {
	CreateUsers(user model.User) error
	GetOneByUsername(username string) (user model.User, err error)
	GetOneByID(id int) (user model.User, err error)
}

type AdapterService interface {
	CreateUserService(user model.User) error
	GetUserByID(id int) (model.User, error)
	GetUserByUsername(username string) (model.User, error)
	LoginUser(username, password string) (string, int)
}
