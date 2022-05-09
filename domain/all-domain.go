package domain

import "github.com/arifwidiasan/todo-app/model"

type AdapterRepository interface {
	CreateUsers(user model.User) error

	GetUserByUsername(username string) (user model.User, err error)
	GetUserByID(id int) (user model.User, err error)

	CreateActivity(activity model.Activity) error
	GetLatestActivity() (activity model.Activity, err error)

	CreateAccess(access model.Access) error
}

type AdapterService interface {
	CheckAuth(username, usernameToken int) error

	CreateUserService(user model.User) error
	GetUserByIDService(id int) (model.User, error)
	GetUserByUsernameService(username string) (model.User, error)
	LoginUser(username, password string) (string, int)

	CreateActivityService(activity model.Activity) error
	GetLatestActivityservice() (model.Activity, error)
	ReplaceActivityName(user_id int, activity model.Activity) model.Activity

	CreateAccessService(user_id, activity_id uint, access model.Access) error
}
