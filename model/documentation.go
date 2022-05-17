package model

type CreateUser struct {
	Username   string `example:"arifwidiasan"`
	User_email string `example:"arifw.subagio17@gmail.com"`
	User_pass  string `example:"rahasia87"`
}

type FailCreateUser struct {
	Messages string `example:"error insert user"`
}

type SuccessLoginUser struct {
	Messages string `example:"success"`
	Token    string `example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTI3ODI1OTIsImlkIjoxLCJ1c2VybmFtZSI6ImFyaWZ3aWRpYXNhbiJ9.Cve64349wPdD4HALRWnLqTzH4ByN20SRT_Zw-Vmrrm4"`
}

type FailLoginUser struct {
	Messages string `example:"email atau password salah"`
}

type InternalError struct {
	Messages string `example:"internal"`
}

type UserFound struct {
	Username   string `gorm:"uniqueIndex;size:30" json:"username" example:"arifwidiasan"`
	User_email string `json:"user_email" example:"arifw.subagio17@gmail.com"`
	User_pass  string `json:"user_pass" example:"rahasia87"`
}

type UserNotFound struct {
	Messages string `example:"username not found"`
}

type JWTNotFound struct {
	Messages string `example:"missing or malformed jwt"`
}

type CreateActivity struct {
	Activity_name string `example:"Kegiatan 1"`
}

type FailCreateActivity struct {
	Messages string `example:"error insert activity"`
}

type ActivityNotFound struct {
	Messages string `example:"activity not found"`
}

type NoAccess struct {
	Messages string `example:"you have no access to this activity"`
}

type NoAccessOwner struct {
	Messages string `example:"you dont have access to this activity or not owner"`
}

type FailUpdateActivity struct {
	Messages string `example:"no id found or no change"`
}

type ActivityDeleted struct {
	Messages string `example:"activity deleted"`
}

type ActivityUpdated struct {
	Messages      string `example:"activity edited"`
	Activity_name string `example:"Kegiatan 1"`
}

type ActivityArchived struct {
	Messages string `example:"activity archived"`
}

type ActivityRestored struct {
	Messages string `example:"activity is restored"`
}

type FailCreateAccess struct {
	Messages string `example:"error insert access"`
}

type CreateAccess struct {
	Username string `example:"arifwidiasan"`
}

type SuccessCreateAccess struct {
	Messages      string `example:"success add another user to this activity"`
	Activity_name string `example:"Kegiatan 1"`
	Username      string `example:"arifwidiasan"`
}

type SuccessDeleteAccess struct {
	Messages      string `example:"success delete user to this activity"`
	Activity_name string `example:"Kegiatan 1"`
	Username      string `example:"arifwidiasan"`
}

type SuccessRemoveAccess struct {
	Messages      string `example:"success delete access, you can't access this activity anymore"`
	Activity_name string `example:"Kegiatan 1"`
}

type AddYourself struct {
	Messages string `example:"can't add yourself"`
}

type FailDeleteAccess struct {
	Messages string `example:"error delete because you're the owner or no access found"`
}

type CreateTask struct {
	Task_name     string `json:"task_name" example:"buat dokumentasi"`
	Task_Priority string `gorm:"type:ENUM('1', '2', '3', '4');default:'1'" json:"task_priority" example:"1"`
}

type FailCreateTask struct {
	Messages string `example:"error insert task"`
}

type SuccessCreateTask struct {
	Messages      string `example:"success add task to this activity"`
	Activity_name string `example:"Kegiatan 1"`
	Task_name     string `example:"buat dokumentasi"`
}

type TaskNotFound struct {
	Messages string `example:"task not found"`
}

type FailUpdateTask struct {
	Messages string `example:"task not found or no change"`
}

type TaskUpdated struct {
	Messages string `example:"success update task"`
	Tasks    Task
}

type TaskDeleted struct {
	Messages string `example:"task deleted"`
}

type TaskCompleted struct {
	Messages string `example:"task completed"`
	Tasks    Task
}

type TaskUndo struct {
	Messages string `example:"success undo completed task"`
	Tasks    Task
}
