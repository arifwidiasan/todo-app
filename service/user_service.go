package service

import (
	"net/http"

	"github.com/arifwidiasan/todo-app/helper"
	"github.com/arifwidiasan/todo-app/model"
)

func (s *svcUser) CreateUserService(user model.User) error {
	return s.repo.CreateUsers(user)
}

func (s *svcUser) GetUserByIDService(id int) (model.User, error) {
	return s.repo.GetUserByID(id)
}

func (s *svcUser) GetUserByUsernameService(username string) (model.User, error) {
	return s.repo.GetUserByUsername(username)
}

func (s *svcUser) LoginUser(username, password string) (string, int) {
	user, _ := s.repo.GetUserByUsername(username)

	if user.User_pass != password {
		return "", http.StatusUnauthorized
	}

	token, err := helper.CreateToken(int(user.ID), user.Username, s.c.JWT_KEY)
	if err != nil {
		return "", http.StatusInternalServerError
	}

	return token, http.StatusOK
}
