package model

type Access struct {
	ID           uint `gorm:"primaryKey;autoIncrement" json:"id"`
	Access_Owner bool `json:"access_owner"`
	UserID       uint `json:"user_id"`
	ActivityID   uint `json:"activity_id"`
}

type ListAccess struct {
	Username     string `json:"username" example:"arifwidiasan"`
	Access_Owner bool   `json:"access_owner" example:"true"`
}
