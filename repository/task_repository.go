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
