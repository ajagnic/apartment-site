package db

// Reservation represents the data model for the application table.
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
