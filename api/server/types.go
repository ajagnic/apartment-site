package server

type Reservation struct {
	Name      string
	Phone     string
	Email     string
	Apartment string
	Guests    string
	Start     string
	End       string
	Date      string `json:"reservationDate"`
}
