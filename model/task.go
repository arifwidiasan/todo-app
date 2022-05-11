package model

import "gorm.io/gorm"

type Task struct {
	*gorm.Model

	Task_Name     string `json:"task_name"`
	Task_Priority string `gorm:"type:ENUM('1', '2', '3', '4');default:'1'" json:"task_priority"`
	Task_Done     bool   `json:"task_done"`
	ActivityID    uint   `json:"activity_id"`
}
