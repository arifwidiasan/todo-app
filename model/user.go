package model

import "gorm.io/gorm"

type User struct {
	*gorm.Model

	Username   string `gorm:"uniqueIndex" json:"username"`
	User_email string `json:"user_email"`
	User_pass  string `json:"user_pass"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
