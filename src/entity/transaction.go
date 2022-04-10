package entity

import (
	"time"
)

type Transaction struct {
	Id        int       `db:"id" json:"id"`
	Event     int       `db:"event_id" json:"event_id"`
	Date      time.Time `db:"transaction_date" json:"transaction_date"`
	Quantity  int       `db:"quantity" json:"quantity"`
	TPrice    int       `db:"total_price" json:"total_price"`
	PayMethod string    `db:"payment_method" json:"payment_method"`
	Action    string    `db:"action" json:"action"`
	Code      int       `db:"code" json:"code"`
	User      []User
}
