package repository

import (
	"github.com/arifwidiasan/todo-app/domain"

	"gorm.io/gorm"
)

type repositoryMysqlLayer struct {
	DB *gorm.DB
}

func NewMysqlRepository(db *gorm.DB) domain.AdapterRepository {
	return &repositoryMysqlLayer{
		DB: db,
	}
}
