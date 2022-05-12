package repository

import (
	"fmt"

	"github.com/arifwidiasan/todo-app/model"
)

func (r *repositoryMysqlLayer) CreateTask(task model.Task) error {
	res := r.DB.Create(&task)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert task")
	}

	return nil
}

func (r *repositoryMysqlLayer) DeleteAllTask(activity_id int) error {
	r.DB.Unscoped().Where("activity_id = ?", activity_id).Delete(&model.Task{})

	return nil
}

func (r *repositoryMysqlLayer) GetAllTask(activity_id int) []model.Task {
	task := []model.Task{}
	r.DB.Model(&model.Task{}).
		Where("activity_id = ?", activity_id).
		Scan(&task)

	return task
}

func (r *repositoryMysqlLayer) GetTaskByID(id int) (task model.Task, err error) {
	res := r.DB.Where("id = ?", id).Find(&task)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("task not found")
	}

	return
}

func (r *repositoryMysqlLayer) UpdateTaskByID(id int, task model.Task) error {
	res := r.DB.Where("id = ?", id).UpdateColumns(&task)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error update task")
	}

	return nil
}

func (r *repositoryMysqlLayer) DeleteTaskByID(id int) error {
	res := r.DB.Unscoped().Delete(&model.Task{}, id)

	if res.RowsAffected < 1 {
		return fmt.Errorf("error delete")
	}

	return nil
}

func (r *repositoryMysqlLayer) CompleteTaskByID(id int, task model.Task) error {
	res := r.DB.Where("id = ?", id).Save(&task)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error update task")
	}

	return nil
}
