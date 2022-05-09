package model

import "gorm.io/gorm"

type Activity struct {
	*gorm.Model

	Activity_Name string   `gorm:"unique" json:"activity_name"`
	Archived      bool     `json:"archived"`
	Accesses      []Access `gorm:"ForeignKey:ActivityID" json:"accesses"`
}