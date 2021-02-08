package db

import "go.mongodb.org/mongo-driver/bson/primitive"

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
	Emailed   bool
}

// Result contains reserved dates for a single record, used by CollectDates.
type Result struct {
	Dates     []string
	Apartment string
}

// Confirmation contains the user email for a record id, used by ReservationsToConfirm.
type Confirmation struct {
	ID    primitive.ObjectID `bson:"_id"`
	Email string
}
