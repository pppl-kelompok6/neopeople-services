package main

import "time"

type Event struct {
	Event_name      string    `json:"event_name"`
	Cover           string    `json:"cover"`
	Date            time.Time `json:"date"`
	Started_at      time.Time `json:"started_at"`
	Finish_at       time.Time `json:"finish_at"`
	Price           int       `json:"price"`
	Speaker         string    `json:"speaker"`
	Speaker_job     string    `json:"speaker_job"`
	Speaker_company string    `json:"speaker_company"`
	Description     string    `json:"description"`
}
