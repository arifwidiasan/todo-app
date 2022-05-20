package service

import (
	"errors"
	"testing"

	"github.com/arifwidiasan/todo-app/config"
	"github.com/arifwidiasan/todo-app/model"
	"github.com/stretchr/testify/assert"
)

func TestCreateActivityService(t *testing.T) {
	testTable := []struct {
		name    string
		f       func(activity model.Activity) error
		noError bool
	}{
		{
			name: "success",
			f: func(activity model.Activity) error {
				return nil
			},
			noError: true,
		},
		{
			name: "error",
			f: func(activity model.Activity) error {
				return errors.New("error")
			},
			noError: false,
		},
	}
	repo := repoMock{}

	for _, v := range testTable {
		t.Run(v.name, func(t *testing.T) {
			repo.createActivity = v.f

			svc := NewServiceUser(&repo, config.Config{})
			err := svc.CreateActivityService(model.Activity{})
			if v.noError {
				assert.NoError(t, err)
			}
		})
	}
}

func TestGetLatestActivityService(t *testing.T) {
	testTable := []struct {
		name    string
		f       func() (activity model.Activity, err error)
		noError bool
	}{
		{
			name: "success",
			f: func() (activity model.Activity, err error) {
				return model.Activity{}, nil
			},
			noError: true,
		},
		{
			name: "error",
			f: func() (activity model.Activity, err error) {
				return model.Activity{}, errors.New("error")
			},
			noError: false,
		},
	}
	repo := repoMock{}

	for _, v := range testTable {
		t.Run(v.name, func(t *testing.T) {
			repo.getLatestActivity = v.f

			svc := NewServiceUser(&repo, config.Config{})
			_, err := svc.GetLatestActivityService()
			if v.noError {
				assert.NoError(t, err)
			}
		})
	}
}

func TestReplaceActivityName(t *testing.T) {
	testTable := []struct {
		name    string
		f       func(user_id int, activity model.Activity) model.Activity
		noError bool
		user_id int
	}{
		{
			name: "success",
			f: func(user_id int, activity model.Activity) model.Activity {
				return activity
			},
			noError: true,
			user_id: 1,
		},
	}
	repo := repoMock{}

	for _, v := range testTable {
		t.Run(v.name, func(t *testing.T) {
			repo.replaceActivityName = v.f

			svc := NewServiceUser(&repo, config.Config{})
			res := svc.ReplaceActivityName(v.user_id, model.Activity{})
			if v.noError {
				assert.Equal(t, res, model.Activity(model.Activity{ID: 0x0, Activity_Name: "-1", Archived: false, Accesses: []model.Access(nil), Tasks: []model.Task(nil)}))
			}
		})
	}
}

func TestGetActivityByNameService(t *testing.T) {
	testTable := []struct {
		name    string
		f       func(name string) (activity model.Activity, err error)
		noError bool
	}{
		{
			name: "success",
			f: func(name string) (activity model.Activity, err error) {
				return model.Activity{}, nil
			},
			noError: true,
		},
		{
			name: "error",
			f: func(name string) (activity model.Activity, err error) {
				return model.Activity{}, errors.New("error")
			},
			noError: false,
		},
	}
	repo := repoMock{}

	for _, v := range testTable {
		t.Run(v.name, func(t *testing.T) {
			repo.getActivityByName = v.f

			svc := NewServiceUser(&repo, config.Config{})
			_, err := svc.GetActivityByNameService("arif")
			if v.noError {
				assert.NoError(t, err)
			}
		})
	}
}

func TestUpdateActivityService(t *testing.T) {
	testTable := []struct {
		name    string
		f       func(id int, activity model.Activity) error
		noError bool
	}{
		{
			name: "success",
			f: func(id int, activity model.Activity) error {
				return nil
			},
			noError: true,
		},
		{
			name: "error",
			f: func(id int, activity model.Activity) error {
				return errors.New("error")
			},
			noError: false,
		},
	}
	repo := repoMock{}

	for _, v := range testTable {
		t.Run(v.name, func(t *testing.T) {
			repo.updateActivityByID = v.f

			svc := NewServiceUser(&repo, config.Config{})
			err := svc.UpdateActivityService(1, model.Activity{})
			if v.noError {
				assert.NoError(t, err)
			}
		})
	}
}

func TestDeleteActivityByIDService(t *testing.T) {
	testTable := []struct {
		name    string
		f       func(id int) error
		noError bool
	}{
		{
			name: "success",
			f: func(id int) error {
				return nil
			},
			noError: true,
		},
		{
			name: "error",
			f: func(id int) error {
				return errors.New("error")
			},
			noError: false,
		},
	}
	repo := repoMock{}

	for _, v := range testTable {
		t.Run(v.name, func(t *testing.T) {
			repo.deleteActivityByID = v.f

			svc := NewServiceUser(&repo, config.Config{})
			err := svc.DeleteActivityByIDService(1)
			if v.noError {
				assert.NoError(t, err)
			}
		})
	}
}

func TestArchiveActivityService(t *testing.T) {
	testTable := []struct {
		name    string
		f       func(id int, activity model.Activity) error
		noError bool
	}{
		{
			name: "success",
			f: func(id int, activity model.Activity) error {
				return nil
			},
			noError: true,
		},
		{
			name: "error",
			f: func(id int, activity model.Activity) error {
				return errors.New("error")
			},
			noError: false,
		},
	}
	repo := repoMock{}

	for _, v := range testTable {
		t.Run(v.name, func(t *testing.T) {
			repo.archiveActivityByID = v.f

			svc := NewServiceUser(&repo, config.Config{})
			err := svc.ArchiveActivityService(1, model.Activity{})
			if v.noError {
				assert.NoError(t, err)
			}
		})
	}
}
