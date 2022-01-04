package model

import "gorm.io/gorm"

type Pantient struct {
	gorm.Model
	Name        string `json:"name" gorm:"type:varchar(100)"`
	Email       string `json:"email" gorm:"type:varchar(100)"`
	PhoneNumber string `json:"phone_number" gorm:"type:varchar(20)"`
	Occupation  string `json:"occupation" gorm:"type:varchar(100)"`
	Company     string `json:"company" gorm:"type:varchar(100)"`
	SessionID   uint
}
