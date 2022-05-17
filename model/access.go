package model

import "gorm.io/gorm"

type Access struct {
	*gorm.Model

	Access_Owner bool `json:"access_owner"`
	UserID       uint `json:"user_id"`
	ActivityID   uint `json:"activity_id"`
}

type ListAccess struct {
	Username     string `example:"arifwidiasan"`
	Access_Owner bool   `example:"true"`
}
