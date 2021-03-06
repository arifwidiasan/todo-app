package model

type User struct {
	ID         uint     `gorm:"primaryKey;autoIncrement" json:"id"`
	Username   string   `gorm:"uniqueIndex;size:30" json:"username"`
	User_email string   `json:"user_email"`
	User_pass  string   `json:"user_pass"`
	Accesses   []Access `gorm:"ForeignKey:UserID" json:"accesses"`
}

type LoginRequest struct {
	Username string `json:"username" example:"arifwidiasan"`
	Password string `json:"password" example:"rahasia87"`
}
