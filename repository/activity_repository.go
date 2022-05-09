package repository

import (
	"fmt"

	"github.com/arifwidiasan/todo-app/model"
)

func (r *repositoryMysqlLayer) CreateActivity(activity model.Activity) error {
	res := r.DB.Create(&activity)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert")
	}

	return nil
}

func (r *repositoryMysqlLayer) GetLatestActivity() (activity model.Activity, err error) {
	res := r.DB.Last(&activity)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}

	return
}
