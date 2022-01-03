package main

import "time"

type session struct {
	Date       time.Time `json:"date"`
	Started_at time.Time `json:"started_at"`
	Finish_at  time.Time `json:"finish_at"`
	Title      string    `json:"title"`
	Platform   string    `json:"platform"`
}
