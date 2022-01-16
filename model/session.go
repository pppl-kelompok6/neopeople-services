package model

import (
	"gorm.io/gorm"
)

type Session struct {
	gorm.Model
	Start     string      `json:"start" gorm:"type:datetime"`
	End       string      `json:"end" gorm:"type:datetime"`
	Title     string      `json:"title" gorm:"type:varchar(100)"`
	Note      string      `json:"note" gorm:"type:varchar(255)"`
	Counselor []Counselor `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Pantient  []Pantient  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
