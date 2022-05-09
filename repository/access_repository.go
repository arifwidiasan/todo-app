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
