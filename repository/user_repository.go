package repository

import (
	"fmt"

	"github.com/arifwidiasan/todo-app/domain"
	"github.com/arifwidiasan/todo-app/model"

	"gorm.io/gorm"
)

type repositoryMysqlLayer struct {
	DB *gorm.DB
}

func (r *repositoryMysqlLayer) CreateUsers(user model.User) error {
	res := r.DB.Create(&user)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert")
	}

	return nil
}

func (r *repositoryMysqlLayer) GetOneByUsername(username string) (user model.User, err error) {
	res := r.DB.Where("username = ?", username).Find(&user)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}

	return
}

func (r *repositoryMysqlLayer) GetOneByID(id int) (user model.User, err error) {
	res := r.DB.Where("id = ?", id).Find(&user)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}

	return
}

func NewMysqlRepository(db *gorm.DB) domain.AdapterRepository {
	return &repositoryMysqlLayer{
		DB: db,
	}
}
