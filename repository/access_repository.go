package repository

import (
	"fmt"

	"github.com/arifwidiasan/todo-app/model"
)

func (r *repositoryMysqlLayer) CreateAccess(access model.Access) error {
	res := r.DB.Create(&access)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert access")
	}

	return nil
}

func (r *repositoryMysqlLayer) CheckAccess(user_id, activity_id uint) (access model.Access, err error) {
	res := r.DB.Where("user_id = ? AND activity_id = ?", user_id, activity_id).Find(&access)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("access not found")
	}

	return
}

func (r *repositoryMysqlLayer) CheckAccessOwner(user_id, activity_id uint) (access model.Access, err error) {
	res := r.DB.Where("access_owner = ? AND user_id = ? AND activity_id = ?", "1", user_id, activity_id).Find(&access)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("access not found, you're not the owner")
	}

	return
}

func (r *repositoryMysqlLayer) DeleteAllAccess(activity_id int) error {
	r.DB.Unscoped().Where("activity_id = ?", activity_id).Delete(&model.Access{})

	return nil
}

func (r *repositoryMysqlLayer) GetAccessUserActivity(activity_id int) []model.ListAccess {
	res := []model.ListAccess{}
	r.DB.Model(&model.User{}).Select("users.username, accesses.access_owner").
		Joins("JOIN accesses on accesses.user_id = users.id").
		Where("accesses.activity_id = ?", activity_id).
		Scan(&res)

	return res
}

func (r *repositoryMysqlLayer) DeleteOneAccess(user_id, activity_id int) error {
	res := r.DB.Unscoped().Where("access_owner = ? AND user_id = ? AND activity_id = ?", "0", user_id, activity_id).Delete(&model.Access{})

	if res.RowsAffected < 1 {
		return fmt.Errorf("error delete because you're the owner or no access found")
	}

	return nil
}
