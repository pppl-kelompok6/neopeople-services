package model

import (
	"gorm.io/gorm"
)

type Session struct {
	gorm.Model
	StartDate string      `json:"startdate" gorm:"type:datetime"`
	EndDate   string      `json:"enddate" gorm:"type:datetime"`
	Title     string      `json:"title" gorm:"type:varchar(100)"`
	Message   string      `json:"message" gorm:"type:varchar(255)"`
	Counselor []Counselor `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Pantient  []Pantient  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
