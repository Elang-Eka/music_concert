package entity

import "time"

type Ticket struct {
	Code      int       `db:"code" json:"code"`
	Event     string    `db:"name" json:"name"`
	Location  string    `db:"location" json:"location"`
	Date      time.Time `db:"date" json:"date"`
	Organizer string    `db:"organizer" json:"organizer"`
}
