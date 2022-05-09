package service

import "github.com/arifwidiasan/todo-app/model"

func (s *svcUser) CreateAccessService(access model.Access) error {
	return s.repo.CreateAccess(access)
}
