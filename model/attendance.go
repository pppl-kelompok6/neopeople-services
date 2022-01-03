package main

type Attendance struct {
	Event_id     int    `json:"event_id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Phome_number string `json:"phone_number"`
	Profession   string `json:"profession"`
	Company      string `json:"company"`
}
