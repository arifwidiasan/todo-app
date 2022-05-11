package service

import "github.com/arifwidiasan/todo-app/model"

func (s *svcUser) CreateTaskService(activity_id uint, task model.Task) error {
	task.Task_Done = false
	task.ActivityID = activity_id

	return s.repo.CreateTask(task)
}

func (s *svcUser) DeleteAllTaskService(activity_id int) error {
	return s.repo.DeleteAllTask(activity_id)
}

func (s *svcUser) GetAllTaskService(activity_id int) []model.Task {
	return s.repo.GetAllTask(activity_id)
}
