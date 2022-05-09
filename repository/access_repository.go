package repository

import (
	"fmt"

	"github.com/arifwidiasan/todo-app/model"
)

func (r *repositoryMysqlLayer) CreateAccess(access model.Access) error {
	res := r.DB.Create(&access)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert access")
	}

	return nil
}

func (r *repositoryMysqlLayer) CheckAccess(user_id, activity_id uint) (access model.Access, err error) {
	res := r.DB.Where("user_id = ? AND activity_id = ?", user_id, activity_id).Find(&access)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("access not found")
	}

	return
}
