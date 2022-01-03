package main

type Event_order struct {
	Event_id       int    `json:"event_id"`
	User_id        int    `json:"user_id"`
	Transaction    string `json:"transaction"`
	Payment_status string `json:"payment_status"`
	Payment_method string `json:"payment_method"`
}
