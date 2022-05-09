package service

import (
	"strconv"
	"strings"

	"github.com/arifwidiasan/todo-app/model"
)

func (s *svcUser) CreateActivityService(activity model.Activity) error {
	return s.repo.CreateActivity(activity)
}

func (s *svcUser) GetLatestActivityservice() (model.Activity, error) {
	return s.repo.GetLatestActivity()
}

func (s *svcUser) ReplaceActivityName(user_id int, activity model.Activity) model.Activity {
	activity.Activity_Name = strings.ReplaceAll(activity.Activity_Name, " ", "-") + "-" + strconv.Itoa(user_id)

	return activity
}
