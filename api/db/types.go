package db

// Reservation represents the form submitted by the ReservationForm component.
type Reservation struct {
	Name      string
	Phone     string
	Email     string
	Apartment string
	Guests    string
	Start     string
	End       string
	Date      string `bson:"reservationDate"`
}
