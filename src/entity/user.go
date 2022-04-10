package entity

type User struct {
	Id     int    `db:"id" json:"id"`
	Name   string `db:"name" json:"name"`
	Age    int    `db:"age" json:"age"`
	Gender string `db:"gender" json:"gender"`
	Email  string `db:"email" json:"email"`
	Id_Trx int    `db:"transaction_id" json:"transaction_id"`
}

type UserTicket struct {
	Id     int    `db:"id" json:"id"`
	Name   string `db:"name" json:"name"`
	Age    int    `db:"age" json:"age"`
	Gender string `db:"gender" json:"gender"`
	Email  string `db:"email" json:"email"`
	Id_Trx int    `db:"transaction_id" json:"transaction_id"`
	Ticket []Ticket
}
