package docs

type CreateUser struct {
	Username   string `example:"arifwidiasan"`
	User_email string `example:"arifw.subagio17@gmail.com"`
	User_pass  string `example:"rahasia87"`
}

type CreateActivity struct {
	Activity_name string `example:"Kegiatan 1"`
}

type CreateAccess struct {
	Username string `example:"arifwidiasan"`
}

type CreateTask struct {
	Task_name     string `json:"task_name" example:"buat dokumentasi"`
	Task_Priority string `gorm:"type:ENUM('1', '2', '3', '4');default:'1'" json:"task_priority" example:"1"`
}
