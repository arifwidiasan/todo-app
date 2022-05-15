package service

import (
	"errors"
	"testing"

	"github.com/arifwidiasan/todo-app/config"
	"github.com/arifwidiasan/todo-app/model"
	"github.com/stretchr/testify/assert"
)

func TestCreateUserService(t *testing.T) {
	testTable := []struct {
		name    string
		f       func(user model.User) error
		noError bool
	}{
		{
			name: "success",
			f: func(user model.User) error {
				return nil
			},
			noError: true,
		},
		{
			name: "error",
			f: func(user model.User) error {
				return errors.New("error")
			},
			noError: false,
		},
	}
	repo := repoMock{}

	for _, v := range testTable {
		t.Run(v.name, func(t *testing.T) {
			repo.createUser = v.f

			svc := NewServiceUser(&repo, config.Config{})
			err := svc.CreateUserService(model.User{})
			if v.noError {
				assert.NoError(t, err)
			}
		})
	}
}

func TestGetUserByUsernameService(t *testing.T) {
	testTable := []struct {
		name    string
		f       func(username string) (model.User, error)
		noError bool
	}{
		{
			name: "success",
			f: func(username string) (model.User, error) {
				return model.User{}, nil
			},
			noError: true,
		},
		{
			name: "error",
			f: func(username string) (model.User, error) {
				return model.User{}, errors.New("error")
			},
			noError: false,
		},
	}
	repo := repoMock{}

	for _, v := range testTable {
		t.Run(v.name, func(t *testing.T) {
			repo.getUserbyUsername = v.f

			svc := NewServiceUser(&repo, config.Config{})
			_, err := svc.GetUserByUsernameService("arif")
			if v.noError {
				assert.NoError(t, err)
			}
		})
	}
}

func TestGetUserByIDervice(t *testing.T) {
	testTable := []struct {
		name    string
		f       func(id int) (model.User, error)
		noError bool
	}{
		{
			name: "success",
			f: func(id int) (model.User, error) {
				return model.User{}, nil
			},
			noError: true,
		},
		{
			name: "error",
			f: func(id int) (model.User, error) {
				return model.User{}, errors.New("error")
			},
			noError: false,
		},
	}
	repo := repoMock{}

	for _, v := range testTable {
		t.Run(v.name, func(t *testing.T) {
			repo.getUserbyID = v.f

			svc := NewServiceUser(&repo, config.Config{})
			_, err := svc.GetUserByIDService(1)
			if v.noError {
				assert.NoError(t, err)
			}
		})
	}
}
