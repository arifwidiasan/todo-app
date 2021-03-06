package repository

import (
	"fmt"

	"github.com/arifwidiasan/todo-app/model"
)

func (r *repositoryMysqlLayer) CreateActivity(activity model.Activity) error {
	res := r.DB.Create(&activity)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert activity")
	}

	return nil
}

func (r *repositoryMysqlLayer) GetLatestActivity() (activity model.Activity, err error) {
	res := r.DB.Last(&activity)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("activity not found")
	}

	return
}

func (r *repositoryMysqlLayer) GetAllActivity(id_user int) []model.Activity {
	activity := []model.Activity{}
	r.DB.Model(&model.Activity{}).Joins("left join accesses on accesses.activity_id = activities.id").
		Where("accesses.user_id = ? AND activities.archived = 0", id_user).
		Scan(&activity)

	return activity
}

func (r *repositoryMysqlLayer) GetAllArchiveActivity(id_user int) []model.Activity {
	activity := []model.Activity{}
	r.DB.Model(&model.Activity{}).Joins("left join accesses on accesses.activity_id = activities.id").
		Where("accesses.user_id = ? AND activities.archived = 1", id_user).
		Scan(&activity)

	return activity
}

func (r *repositoryMysqlLayer) GetActivityByName(name string) (activity model.Activity, err error) {
	res := r.DB.Where("activity_name = ?", name).Find(&activity)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("activity not found")
	}

	return
}

func (r *repositoryMysqlLayer) UpdateActivityByID(id int, activity model.Activity) error {
	res := r.DB.Where("id = ?", id).UpdateColumns(&activity)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error update activity")
	}

	return nil
}

func (r *repositoryMysqlLayer) ArchiveActivityByID(id int, activity model.Activity) error {
	res := r.DB.Where("id = ?", id).Save(&activity)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error update activity")
	}

	return nil
}

func (r *repositoryMysqlLayer) DeleteActivityByID(id int) error {
	res := r.DB.Unscoped().Delete(&model.Activity{}, id)

	if res.RowsAffected < 1 {
		return fmt.Errorf("error delete")
	}

	return nil
}
