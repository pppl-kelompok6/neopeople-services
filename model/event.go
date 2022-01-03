package model

import "gorm.io/gorm"

type Event struct {
	gorm.Model
	EventName      string       `json:"event_name" gorm:"type:varchar(100)"`
	Cover          string       `json:"cover" gorm:"type:varchar(255)"`
	Date           string       `json:"date" gorm:"type:date"`
	StartedAt      string       `json:"started_at" gorm:"type:datetime"`
	FinishAt       string       `json:"finish_at" gorm:"type:datetime"`
	Price          int          `json:"price"`
	Speaker        string       `json:"speaker" gorm:"type:varchar(100)"`
	SpeakerJob     string       `json:"speaker_job" gorm:"type:varchar(100)"`
	SpeakerCompany string       `json:"speaker_company" gorm:"type:varchar(100)"`
	Description    string       `json:"description" gorm:"type:varchar(255)"`
	Attendance     []Attendance `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	EventOrder     []EventOrder `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
