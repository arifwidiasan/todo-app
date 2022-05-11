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
		Where("activity_id = ? AND task_done = 0", activity_id).
		Scan(&task)

	return task
}
