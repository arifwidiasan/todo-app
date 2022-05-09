package service

import (
	"strconv"
	"strings"

	"github.com/arifwidiasan/todo-app/model"
)

func (s *svcUser) CreateActivityService(activity model.Activity) error {
	return s.repo.CreateActivity(activity)
}

func (s *svcUser) GetLatestActivityService() (model.Activity, error) {
	return s.repo.GetLatestActivity()
}

func (s *svcUser) GetAllActivityService(username string) []model.Activity {
	user, _ := s.repo.GetUserByUsername(username)
	return s.repo.GetAllActivity(int(user.ID))
}

func (s *svcUser) ReplaceActivityName(user_id int, activity model.Activity) model.Activity {
	activity.Activity_Name = strings.ReplaceAll(activity.Activity_Name, " ", "-") + "-" + strconv.Itoa(user_id)

	return activity
}
