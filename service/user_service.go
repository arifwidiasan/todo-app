package service

import (
	"net/http"

	"github.com/arifwidiasan/todo-app/config"
	"github.com/arifwidiasan/todo-app/domain"
	"github.com/arifwidiasan/todo-app/helper"
	"github.com/arifwidiasan/todo-app/model"
)

type svcUser struct {
	c    config.Config
	repo domain.AdapterRepository
}

func (s *svcUser) CreateUserService(user model.User) error {
	return s.repo.CreateUsers(user)
}

func (s *svcUser) GetUserByID(id int) (model.User, error) {
	return s.repo.GetOneByID(id)
}

func (s *svcUser) GetUserByUsername(username string) (model.User, error) {
	return s.repo.GetOneByUsername(username)
}

func (s *svcUser) LoginUser(username, password string) (string, int) {
	user, _ := s.repo.GetOneByUsername(username)

	if (user.User_pass != password) || (user == model.User{}) {
		return "", http.StatusUnauthorized
	}

	token, err := helper.CreateToken(int(user.ID), user.Username, s.c.JWT_KEY)
	if err != nil {
		return "", http.StatusInternalServerError
	}

	return token, http.StatusOK
}

func NewServiceUser(repo domain.AdapterRepository, c config.Config) domain.AdapterService {
	return &svcUser{
		repo: repo,
		c:    c,
	}
}
