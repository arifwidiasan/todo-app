package service

import (
	"errors"

	"github.com/arifwidiasan/todo-app/model"
)

func (s *svcUser) CreateAccessService(user_id, activity_id uint, owner bool, access model.Access) error {
	access.Access_Owner = owner
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

func (s *svcUser) CheckAcccessOwnerService(user_id, activity_id uint) error {
	_, err := s.repo.CheckAccessOwner(user_id, activity_id)
	if err != nil {
		return errors.New("error")
	}

	return nil
}

func (s *svcUser) DeleteAllAccessService(activity_id int) error {
	return s.repo.DeleteAllAccess(activity_id)
}

func (s *svcUser) GetAccessUserActivityService(activity_id int) []model.ListAccess {
	return s.repo.GetAccessUserActivity(activity_id)
}

func (s *svcUser) DeleteOneAccessService(user_id, activity_id int) error {
	return s.repo.DeleteOneAccess(user_id, activity_id)
}
