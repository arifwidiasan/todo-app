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

func (s *svcUser) GetTaskByIDService(id int) (model.Task, error) {
	return s.repo.GetTaskByID(id)
}

func (s *svcUser) UpdateTaskService(id int, task model.Task) error {
	return s.repo.UpdateTaskByID(id, task)
}

func (s *svcUser) DeleteTaskByIDService(id int) error {
	return s.repo.DeleteTaskByID(id)
}

func (s *svcUser) CompleteTaskService(id int, task model.Task) error {
	return s.repo.CompleteTaskByID(id, task)
}
