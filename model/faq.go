package model

import "gorm.io/gorm"

type Faq struct {
	gorm.Model
	Category string `json:"category" gorm:"type:varchar(10)"`
	Question string `json:"question" gorm:"type:varchar(255)"`
	Answer   string `json:"answer" gorm:"type:varchar(255)"`
}
