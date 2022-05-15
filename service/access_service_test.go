package service

import (
	"errors"
	"testing"

	"github.com/arifwidiasan/todo-app/config"
	"github.com/arifwidiasan/todo-app/model"
	"github.com/stretchr/testify/assert"
	//"github.com/stretchr/testify/mock"
)

func TestCreateAccessService(t *testing.T) {
	testTable := []struct {
		name    string
		f       func(access model.Access) error
		noError bool
	}{
		{
			name: "success",
			f: func(access model.Access) error {
				return nil
			},
			noError: true,
		},
		{
			name: "error",
			f: func(access model.Access) error {
				return errors.New("error")
			},
			noError: false,
		},
	}
	repo := repoMock{}

	for _, v := range testTable {
		t.Run(v.name, func(t *testing.T) {
			repo.createAccess = v.f

			svc := NewServiceUser(&repo, config.Config{})
			err := svc.CreateAccessService(1, 1, true, model.Access{})
			if v.noError {
				assert.NoError(t, err)
			}
		})
	}
}

func TestCheckAccessService(t *testing.T) {
	testTable := []struct {
		name    string
		f       func(user_id, activity_id uint) (access model.Access, err error)
		noError bool
	}{
		{
			name: "success",
			f: func(user_id, activity_id uint) (access model.Access, err error) {
				return access, nil
			},
			noError: true,
		},
		{
			name: "error",
			f: func(user_id, activity_id uint) (access model.Access, err error) {
				return access, errors.New("error")
			},
			noError: false,
		},
	}
	repo := repoMock{}

	for _, v := range testTable {
		t.Run(v.name, func(t *testing.T) {
			repo.checkAccess = v.f

			svc := NewServiceUser(&repo, config.Config{})
			err := svc.CheckAcccessService(1, 1)
			if v.noError {
				assert.NoError(t, err)
			}
		})
	}
}

func TestCheckAccessOwnerService(t *testing.T) {
	testTable := []struct {
		name    string
		f       func(user_id, activity_id uint) (access model.Access, err error)
		noError bool
	}{
		{
			name: "success",
			f: func(user_id, activity_id uint) (access model.Access, err error) {
				return access, nil
			},
			noError: true,
		},
		{
			name: "error",
			f: func(user_id, activity_id uint) (access model.Access, err error) {
				return access, errors.New("error")
			},
			noError: false,
		},
	}
	repo := repoMock{}

	for _, v := range testTable {
		t.Run(v.name, func(t *testing.T) {
			repo.checkAccessOwner = v.f

			svc := NewServiceUser(&repo, config.Config{})
			err := svc.CheckAcccessOwnerService(1, 1)
			if v.noError {
				assert.NoError(t, err)
			}
		})
	}
}

func TestDeleteAllAccessService(t *testing.T) {
	testTable := []struct {
		name    string
		f       func(activity_id int) error
		noError bool
	}{
		{
			name: "success",
			f: func(activity_id int) error {
				return nil
			},
			noError: true,
		},
		{
			name: "error",
			f: func(activity_id int) error {
				return errors.New("error")
			},
			noError: false,
		},
	}
	repo := repoMock{}

	for _, v := range testTable {
		t.Run(v.name, func(t *testing.T) {
			repo.deleteAllAccess = v.f

			svc := NewServiceUser(&repo, config.Config{})
			err := svc.DeleteAllAccessService(1)
			if v.noError {
				assert.NoError(t, err)
			}
		})
	}
}

func TestGetAccessUserActivityService(t *testing.T) {
	testTable := []struct {
		name    string
		f       func(activity_id int) []model.ListAccess
		noError bool
	}{
		{
			name: "success",
			f: func(activity_id int) []model.ListAccess {
				return []model.ListAccess{}
			},
			noError: true,
		},
		{
			name: "error",
			f: func(activity_id int) []model.ListAccess {
				return []model.ListAccess{}
			},
			noError: false,
		},
	}
	repo := repoMock{}

	for _, v := range testTable {
		t.Run(v.name, func(t *testing.T) {
			repo.getAccessUserActivity = v.f

			svc := NewServiceUser(&repo, config.Config{})
			res := svc.GetAccessUserActivityService(1)
			if v.noError {
				assert.Equal(t, res, []model.ListAccess([]model.ListAccess{}))
			}
		})
	}
}

func TestDeleteOneAccessService(t *testing.T) {
	testTable := []struct {
		name    string
		f       func(user_id, activity_id int) error
		noError bool
	}{
		{
			name: "success",
			f: func(user_id, activity_id int) error {
				return nil
			},
			noError: true,
		},
		{
			name: "error",
			f: func(user_id, activity_id int) error {
				return errors.New("error")
			},
			noError: false,
		},
	}
	repo := repoMock{}

	for _, v := range testTable {
		t.Run(v.name, func(t *testing.T) {
			repo.deleteOneAccess = v.f

			svc := NewServiceUser(&repo, config.Config{})
			err := svc.DeleteOneAccessService(1, 1)
			if v.noError {
				assert.NoError(t, err)
			}
		})
	}
}
