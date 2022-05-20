package model

type Activity struct {
	ID            uint     `gorm:"primaryKey;autoIncrement" json:"id"`
	Activity_Name string   `gorm:"unique" json:"activity_name"`
	Archived      bool     `json:"archived"`
	Accesses      []Access `gorm:"ForeignKey:ActivityID" json:"accesses"`
	Tasks         []Task   `gorm:"ForeignKey:ActivityID" json:"tasks"`
}

func (a *Activity) Archive() {
	a.Archived = true
}

func (a *Activity) Restore() {
	a.Archived = false
}
