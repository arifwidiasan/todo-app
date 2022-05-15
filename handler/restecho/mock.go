package handler

import (
	"github.com/arifwidiasan/todo-app/model"
	"github.com/stretchr/testify/mock"
)

type MockSvc struct {
	mock.Mock
}

func (m *MockSvc) CheckAuth(id, idToken int) error {
	return nil
}

func (m *MockSvc) CreateUserService(user model.User) error {
	ret := m.Called()

	return ret.Error(0)
}
func (m *MockSvc) GetUserByIDService(id int) (model.User, error) {
	return model.User{}, nil
}
func (m *MockSvc) GetUserByUsernameService(username string) (model.User, error) {
	ret := m.Called()

	return model.User{}, ret.Error(0)
}
func (m *MockSvc) LoginUser(username, password string) (string, int) {
	return "success", 200
}

func (m *MockSvc) CreateActivityService(activity model.Activity) error {
	ret := m.Called()

	return ret.Error(0)
}
func (m *MockSvc) GetLatestActivityService() (model.Activity, error) {
	return model.Activity{}, nil
}
func (m *MockSvc) ReplaceActivityName(user_id int, activity model.Activity) model.Activity {
	return model.Activity{}
}
func (m *MockSvc) GetAllActivityService(username string) []model.Activity {
	return []model.Activity{}
}
func (m *MockSvc) GetAllArchiveActivityService(username string) []model.Activity {
	return []model.Activity{}
}
func (m *MockSvc) GetActivityByNameService(name string) (model.Activity, error) {
	return model.Activity{}, nil
}
func (m *MockSvc) UpdateActivityService(id int, activity model.Activity) error {
	return nil
}
func (m *MockSvc) DeleteActivityByIDService(id int) error {
	return nil
}
func (m *MockSvc) ArchiveActivityService(id int, activity model.Activity) error {
	return nil
}

func (m *MockSvc) CreateAccessService(user_id, activity_id uint, owner bool, access model.Access) error {
	ret := m.Called()

	return ret.Error(0)
}
func (m *MockSvc) CheckAcccessService(user_id, activity_id uint) error {
	return nil
}
func (m *MockSvc) CheckAcccessOwnerService(user_id, activity_id uint) error {
	return nil
}
func (m *MockSvc) DeleteAllAccessService(activity_id int) error {
	return nil
}
func (m *MockSvc) GetAccessUserActivityService(activity_id int) []model.ListAccess {
	return []model.ListAccess{}
}
func (m *MockSvc) DeleteOneAccessService(user_id, activity_id int) error {
	return nil
}

func (m *MockSvc) CreateTaskService(activity_id uint, task model.Task) error {
	ret := m.Called()

	return ret.Error(0)
}
func (m *MockSvc) DeleteAllTaskService(activity_id int) error {
	return nil
}
func (m *MockSvc) GetAllTaskService(activity_id int) []model.Task {
	return []model.Task{}
}
func (m *MockSvc) GetTaskByIDService(id int) (model.Task, error) {
	return model.Task{}, nil
}
func (m *MockSvc) UpdateTaskService(id int, task model.Task) error {
	return nil
}
func (m *MockSvc) DeleteTaskByIDService(id int) error {
	return nil
}
func (m *MockSvc) CompleteTaskService(id int, task model.Task) error {
	return nil
}
