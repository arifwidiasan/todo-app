package service

import "github.com/arifwidiasan/todo-app/model"

func (s *svcUser) CreateTaskService(activity_id uint, task model.Task) error {
	task.Task_Done = false
	task.ActivityID = activity_id

	return s.repo.CreateTask(task)
}
