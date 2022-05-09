package service

import (
	"errors"

	"github.com/arifwidiasan/todo-app/config"
	"github.com/arifwidiasan/todo-app/domain"
)

type svcUser struct {
	c    config.Config
	repo domain.AdapterRepository
}

func (s *svcUser) CheckAuth(username, usernameToken int) error {
	if username != usernameToken {
		return errors.New("error")
	}

	return nil
}

func NewServiceUser(repo domain.AdapterRepository, c config.Config) domain.AdapterService {
	return &svcUser{
		repo: repo,
		c:    c,
	}
}
