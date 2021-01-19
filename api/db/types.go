package db

// Reservation represents the form submitted by the ReservationForm component.
type Reservation struct {
	Name      string
	Phone     string
	Email     string
	Apartment string
	Guests    int
	Dates     []string
	Created   string
	Confirmed bool
	Cancelled bool
}

// Result contains reserved dates for a single record, used by CollectDates.
type Result struct {
	Dates     []string
	Apartment string
}
