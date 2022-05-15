package service

import (
	"errors"
	"testing"

	"github.com/arifwidiasan/todo-app/config"
	"github.com/arifwidiasan/todo-app/model"
	"github.com/stretchr/testify/assert"
	//"github.com/stretchr/testify/mock"
)

func TestCreateTaskService(t *testing.T) {
	testTable := []struct {
		name    string
		f       func(task model.Task) error
		noError bool
	}{
		{
			name: "success",
			f: func(task model.Task) error {
				return nil
			},
			noError: true,
		},
		{
			name: "error",
			f: func(task model.Task) error {
				return errors.New("error")
			},
			noError: false,
		},
	}
	repo := repoMock{}

	for _, v := range testTable {
		t.Run(v.name, func(t *testing.T) {
			repo.createTask = v.f

			svc := NewServiceUser(&repo, config.Config{})
			err := svc.CreateTaskService(1, model.Task{})
			if v.noError {
				assert.NoError(t, err)
			}
		})
	}
}

func TestDeleteAllTaskService(t *testing.T) {
	testTable := []struct {
		name    string
		f       func(activity_id int) error
		noError bool
	}{
		{
			name: "success",
			f: func(activity_id int) error {
				return nil
			},
			noError: true,
		},
		{
			name: "error",
			f: func(activity_id int) error {
				return errors.New("error")
			},
			noError: false,
		},
	}
	repo := repoMock{}

	for _, v := range testTable {
		t.Run(v.name, func(t *testing.T) {
			repo.deleteAllTask = v.f

			svc := NewServiceUser(&repo, config.Config{})
			err := svc.DeleteAllTaskService(1)
			if v.noError {
				assert.NoError(t, err)
			}
		})
	}
}

func TestGetAllTaskService(t *testing.T) {
	testTable := []struct {
		name    string
		f       func(activity_id int) []model.Task
		noError bool
	}{
		{
			name: "success",
			f: func(activity_id int) []model.Task {
				return []model.Task{}
			},
			noError: true,
		},
		{
			name: "error",
			f: func(activity_id int) []model.Task {
				return []model.Task{}
			},
			noError: false,
		},
	}
	repo := repoMock{}

	for _, v := range testTable {
		t.Run(v.name, func(t *testing.T) {
			repo.getAllTask = v.f

			svc := NewServiceUser(&repo, config.Config{})
			res := svc.GetAllTaskService(1)
			if v.noError {
				assert.Equal(t, res, []model.Task{})
			}
		})
	}
}

func TestGetTaskByIDService(t *testing.T) {
	testTable := []struct {
		name    string
		f       func(id int) (task model.Task, err error)
		noError bool
	}{
		{
			name: "success",
			f: func(id int) (task model.Task, err error) {
				return task, nil
			},
			noError: true,
		},
		{
			name: "error",
			f: func(id int) (task model.Task, err error) {
				return task, errors.New("error")
			},
			noError: false,
		},
	}
	repo := repoMock{}

	for _, v := range testTable {
		t.Run(v.name, func(t *testing.T) {
			repo.getTaskByID = v.f

			svc := NewServiceUser(&repo, config.Config{})
			_, err := svc.GetTaskByIDService(1)
			if v.noError {
				assert.NoError(t, err)
			}
		})
	}
}

func TestUpdateTaskService(t *testing.T) {
	testTable := []struct {
		name    string
		f       func(id int, task model.Task) error
		noError bool
	}{
		{
			name: "success",
			f: func(id int, task model.Task) error {
				return nil
			},
			noError: true,
		},
		{
			name: "error",
			f: func(id int, task model.Task) error {
				return errors.New("error")
			},
			noError: false,
		},
	}
	repo := repoMock{}

	for _, v := range testTable {
		t.Run(v.name, func(t *testing.T) {
			repo.updateTaskByID = v.f

			svc := NewServiceUser(&repo, config.Config{})
			err := svc.UpdateTaskService(1, model.Task{})
			if v.noError {
				assert.NoError(t, err)
			}
		})
	}
}

func TestDeleteTaskByIDService(t *testing.T) {
	testTable := []struct {
		name    string
		f       func(id int) error
		noError bool
	}{
		{
			name: "success",
			f: func(id int) error {
				return nil
			},
			noError: true,
		},
		{
			name: "error",
			f: func(id int) error {
				return errors.New("error")
			},
			noError: false,
		},
	}
	repo := repoMock{}

	for _, v := range testTable {
		t.Run(v.name, func(t *testing.T) {
			repo.deleteTaskByID = v.f

			svc := NewServiceUser(&repo, config.Config{})
			err := svc.DeleteTaskByIDService(1)
			if v.noError {
				assert.NoError(t, err)
			}
		})
	}
}

func TestCompleteTaskService(t *testing.T) {
	testTable := []struct {
		name    string
		f       func(id int, task model.Task) error
		noError bool
	}{
		{
			name: "success",
			f: func(id int, task model.Task) error {
				return nil
			},
			noError: true,
		},
		{
			name: "error",
			f: func(id int, task model.Task) error {
				return errors.New("error")
			},
			noError: false,
		},
	}
	repo := repoMock{}

	for _, v := range testTable {
		t.Run(v.name, func(t *testing.T) {
			repo.completeTaskByID = v.f

			svc := NewServiceUser(&repo, config.Config{})
			err := svc.CompleteTaskService(1, model.Task{})
			if v.noError {
				assert.NoError(t, err)
			}
		})
	}
}
