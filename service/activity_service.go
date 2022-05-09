package service

import "github.com/arifwidiasan/todo-app/model"

func (s *svcUser) CreateActivityService(activity model.Activity) error {
	return s.repo.CreateActivity(activity)
}

func (s *svcUser) GetLatestActivityservice() (model.Activity, error) {
	return s.repo.GetLatestActivity()
}
