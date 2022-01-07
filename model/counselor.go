package model

import "gorm.io/gorm"

type Counselor struct {
	gorm.Model
	Username    string `json:"username" gorm:"type:varchar(100);unique_index"`
	Password    string `json:"password" gorm:"type:varchar(100)"`
	Email       string `json:"email" gorm:"type:varchar(100);unique_index"`
	Name        string `json:"name"  gorm:"type:varchar(100)"`
	PhoneNumber string `json:"phone_number" gorm:"type:varchar(20);unique_index"`
	Occupation  string `json:"occupation" gorm:"type:varchar(100)"`
	Company     string `json:"company" gorm:"type:varchar(100)"`
	Position    string `json:"position" gorm:"type:ENUM('CEO', 'Founder', 'Mentor', 'PublicRelation','Team')"`
	Photo       string `json:"photo" gorm:"type:varchar(255)"`
	SessionID   uint   `json:"session_id"`
}
