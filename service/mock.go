package service

import "github.com/arifwidiasan/todo-app/model"

type repoMock struct {
	createUser            func(user model.User) error
	getUserbyUsername     func(username string) (user model.User, err error)
	getUserbyID           func(id int) (user model.User, err error)
	createActivity        func(activity model.Activity) error
	getLatestActivity     func() (activity model.Activity, err error)
	replaceActivityName   func(user_id int, activity model.Activity) model.Activity
	getActivityByName     func(name string) (activity model.Activity, err error)
	updateActivityByID    func(id int, activity model.Activity) error
	deleteActivityByID    func(id int) error
	archiveActivityByID   func(id int, activity model.Activity) error
	createAccess          func(access model.Access) error
	checkAccess           func(user_id, activity_id uint) (access model.Access, err error)
	checkAccessOwner      func(user_id, activity_id uint) (access model.Access, err error)
	deleteAllAccess       func(activity_id int) error
	getAccessUserActivity func(activity_id int) []model.ListAccess
	deleteOneAccess       func(user_id, activity_id int) error
	createTask            func(task model.Task) error
	deleteAllTask         func(activity_id int) error
	getAllTask            func(activity_id int) []model.Task
	getTaskByID           func(id int) (task model.Task, err error)
	updateTaskByID        func(id int, task model.Task) error
	deleteTaskByID        func(id int) error
	completeTaskByID      func(id int, task model.Task) error
	getAllActivity        func(id_user int) []model.Activity
}

func (r *repoMock) CreateUsers(user model.User) error {
	return r.createUser(user)
}
func (r *repoMock) GetUserByUsername(username string) (user model.User, err error) {
	return r.getUserbyUsername(username)
}
func (r *repoMock) GetUserByID(id int) (user model.User, err error) {
	return r.getUserbyID(id)
}

func (r *repoMock) CreateActivity(activity model.Activity) error {
	return r.createActivity(activity)
}
func (r *repoMock) GetLatestActivity() (activity model.Activity, err error) {
	return r.getLatestActivity()
}
func (r *repoMock) GetAllActivity(id_user int) []model.Activity {
	return r.getAllActivity(id_user)
}
func (r *repoMock) GetAllArchiveActivity(id_user int) []model.Activity {
	panic("impl")
}
func (r *repoMock) GetActivityByName(name string) (activity model.Activity, err error) {
	return r.getActivityByName(name)
}
func (r *repoMock) UpdateActivityByID(id int, activity model.Activity) error {
	return r.updateActivityByID(id, activity)
}
func (r *repoMock) DeleteActivityByID(id int) error {
	return r.deleteActivityByID(id)
}
func (r *repoMock) ArchiveActivityByID(id int, activity model.Activity) error {
	return r.archiveActivityByID(id, activity)
}

func (r *repoMock) CreateAccess(access model.Access) error {
	return r.createAccess(access)
}
func (r *repoMock) CheckAccess(user_id, activity_id uint) (access model.Access, err error) {
	return r.checkAccess(user_id, activity_id)
}
func (r *repoMock) CheckAccessOwner(user_id, activity_id uint) (access model.Access, err error) {
	return r.checkAccessOwner(user_id, activity_id)
}
func (r *repoMock) DeleteAllAccess(activity_id int) error {
	return r.deleteAllAccess(activity_id)
}
func (r *repoMock) GetAccessUserActivity(activity_id int) []model.ListAccess {
	return r.getAccessUserActivity(activity_id)
}
func (r *repoMock) DeleteOneAccess(user_id, activity_id int) error {
	return r.deleteOneAccess(user_id, activity_id)
}

func (r *repoMock) CreateTask(task model.Task) error {
	return r.createTask(task)
}
func (r *repoMock) DeleteAllTask(activity_id int) error {
	return r.deleteAllTask(activity_id)
}
func (r *repoMock) GetAllTask(activity_id int) []model.Task {
	return r.getAllTask(activity_id)
}
func (r *repoMock) GetTaskByID(id int) (task model.Task, err error) {
	return r.getTaskByID(id)
}
func (r *repoMock) UpdateTaskByID(id int, task model.Task) error {
	return r.updateTaskByID(id, task)
}
func (r *repoMock) DeleteTaskByID(id int) error {
	return r.deleteTaskByID(id)
}
func (r *repoMock) CompleteTaskByID(id int, task model.Task) error {
	return r.completeTaskByID(id, task)
}
