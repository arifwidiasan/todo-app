package service

import (
	"errors"

	"github.com/arifwidiasan/todo-app/model"
)

func (s *svcUser) CreateAccessService(user_id, activity_id uint, access model.Access) error {
	access.Access_Owner = true
	access.UserID = user_id
	access.ActivityID = activity_id

	return s.repo.CreateAccess(access)
}

func (s *svcUser) CheckAcccessService(user_id, activity_id uint) error {
	_, err := s.repo.CheckAccess(user_id, activity_id)
	if err != nil {
		return errors.New("error")
	}

	return nil
}
