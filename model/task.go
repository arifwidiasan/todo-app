package model

type Task struct {
	ID            uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Task_Name     string `json:"task_name"`
	Task_Priority string `gorm:"type:ENUM('1', '2', '3', '4');default:'1'" json:"task_priority"`
	Task_Done     bool   `json:"task_done"`
	ActivityID    uint   `json:"activity_id"`
}

func (t *Task) Complete() {
	t.Task_Done = true
}

func (t *Task) Undo() {
	t.Task_Done = false
}

type Task_Priority struct {
	ID                 int
	Task_Priority_Name string
}
