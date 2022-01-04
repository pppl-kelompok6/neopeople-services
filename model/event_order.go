package model

type EventOrder struct {
	MidtransOrderID string `json:"midtrans_order_id"`
	Payment_status  string `json:"payment_status" gorm:"type:ENUM('Success','Failed', 'Pending')"`
	Payment_method  string `json:"payment_method" gorm:"type:ENUM('Dana', 'Gopay', 'Ovo')"`
	Amount          int    `json:"amount"`
	EventId         uint
}
