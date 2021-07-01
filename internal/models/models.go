package models

import (
	"time"
)

// User is the user model
type User struct {
	ID 		 	int
	FirstName 	string
	LastName  	string
	Email		 	string
	Password  	string
	AccessLevel int
	Created		time.Time
	Updated 		time.Time
}
// Room is the room model
type Room struct {
	ID 			int
	RoomName 	string
	Created		time.Time
	Updated 		time.Time
}
// Restriction is the restriction model
type Restriction struct {
	ID 					int
	RestrictionName 	string
	Created				time.Time
	Updated 				time.Time
}
// Reservations is the reservation model
type Reservation struct {
	ID 					int
	FirstName 			string
	LastName  			string
	Email		 			string
	Phone					string
	StartDate			time.Time
	EndDate				time.Time
	RoomID				int
	Created				time.Time
	Updated 				time.Time
	Room 					Room
	Processed			int
}

// RoomRestrictions is the roomRestrictions model
type RoomRestriction struct {
	ID 					int
	StartDate			time.Time
	EndDate				time.Time
	RoomID				int
	ReservationID		int
	RestrictionID		int
	Created				time.Time
	Updated 				time.Time
	Room 					Room
	Reservation			Reservation
	Restriction			Restriction
}

// MailData holds an email message
type MailData struct {
	To 		string
	From 		string
	Subject 	string
	Content	string
	Template string
}
