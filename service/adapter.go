package service

import (
	"github.com/arifwidiasan/todo-app/config"
	"github.com/arifwidiasan/todo-app/domain"
)

type svcUser struct {
	c    config.Config
	repo domain.AdapterRepository
}

func NewServiceUser(repo domain.AdapterRepository, c config.Config) domain.AdapterService {
	return &svcUser{
		repo: repo,
		c:    c,
	}
}
