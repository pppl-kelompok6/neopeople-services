package main

import "time"

type counselor_id struct {
	User_id      int       `json:"user_id"`
	Counselor_id int       `json:"counselor_id"`
	Date         time.Time `json:"date"`
	Started_at   time.Time `json:"started_at"`
	Finish_at    time.Time `json:"finish_at"`
	Title        string    `json:"title"`
	Platform     string    `json:"platform"`
}
