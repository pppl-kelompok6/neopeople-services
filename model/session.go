package model

import (
	"gorm.io/gorm"
)

type Session struct {
	gorm.Model
	Date      string      `json:"date" gorm:"type:date"`
	StartedAt string      `json:"started_at" gorm:"type:datetime"`
	FinishAt  string      `json:"finish_at" gorm:"type:datetime"`
	Title     string      `json:"title" gorm:"type:varchar(100)"`
	Platform  string      `json:"platform" gorm:"type:varchar(255)"`
	Counselor []Counselor `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Pantient  []Pantient  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
