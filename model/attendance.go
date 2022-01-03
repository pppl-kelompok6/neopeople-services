package model

import "gorm.io/gorm"

type Attendance struct {
	gorm.Model
	Name         string `json:"name" gorm:"type:varchar(100)"`
	Email        string `json:"email" gorm:"type:varchar(100);unique_index"`
	Phome_number string `json:"phone_number" gorm:"type:varchar(20);unique_index"`
	Profession   string `json:"profession" gorm:"type:varchar(100)"`
	Company      string `json:"company" gorm:"type:varchar(100)"`
}
