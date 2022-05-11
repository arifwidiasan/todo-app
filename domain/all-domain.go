package domain

import "github.com/arifwidiasan/todo-app/model"

type AdapterRepository interface {
	CreateUsers(user model.User) error

	GetUserByUsername(username string) (user model.User, err error)
	GetUserByID(id int) (user model.User, err error)

	CreateActivity(activity model.Activity) error
	GetLatestActivity() (activity model.Activity, err error)
	GetAllActivity(id_user int) []model.Activity
	GetAllArchiveActivity(id_user int) []model.Activity
	GetActivityByName(name string) (activity model.Activity, err error)
	UpdateActivityByID(id int, activity model.Activity) error
	DeleteActivityByID(id int) error
	ArchiveActivityByID(id int, activity model.Activity) error

	CreateAccess(access model.Access) error
	CheckAccess(user_id, activity_id uint) (access model.Access, err error)
	CheckAccessOwner(user_id, activity_id uint) (access model.Access, err error)
	DeleteAllAccess(activity_id int) error
	GetAccessUserActivity(activity_id int) []model.ListAccess
	DeleteOneAccess(user_id, activity_id int) error

	CreateTask(task model.Task) error
}

type AdapterService interface {
	CheckAuth(id, idToken int) error

	CreateUserService(user model.User) error
	GetUserByIDService(id int) (model.User, error)
	GetUserByUsernameService(username string) (model.User, error)
	LoginUser(username, password string) (string, int)

	CreateActivityService(activity model.Activity) error
	GetLatestActivityService() (model.Activity, error)
	ReplaceActivityName(user_id int, activity model.Activity) model.Activity
	GetAllActivityService(username string) []model.Activity
	GetAllArchiveActivityService(username string) []model.Activity
	GetActivityByNameService(name string) (model.Activity, error)
	UpdateActivityService(id int, activity model.Activity) error
	DeleteActivityByIDService(id int) error
	ArchiveActivityService(id int, activity model.Activity) error

	CreateAccessService(user_id, activity_id uint, owner bool, access model.Access) error
	CheckAcccessService(user_id, activity_id uint) error
	CheckAcccessOwnerService(user_id, activity_id uint) error
	DeleteAllAccessService(activity_id int) error
	GetAccessUserActivityService(activity_id int) []model.ListAccess
	DeleteOneAccessService(user_id, activity_id int) error

	CreateTaskService(activity_id uint, task model.Task) error
}
