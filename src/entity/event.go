package entity

// A User belong to the domain layer.
type Event struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Price     int    `json:"price"`
	Location  string `json:"location"`
	Date      string `json:"date"`
	Time      string `json:"time"`
	Organizer string `json:"organizer"`
}
