package models

import (
	"time"
)

// User is the user model
type User struct {
	ID 		 	int			`json:"id"`
	FirstName 	string		`json:"first_name"`
	LastName  	string		`json:"last_name"`
	Email		 	string		`json:"email"`
	Password  	string		`json:"password"`
	AccessLevel int			`json:"access_level"`
	Created		time.Time	`json:"created_at"`
	Updated 		time.Time	`json:"updated_at"`
}
// Room is the room model
type Room struct {
	ID 			int			`json:"id"`
	RoomName 	string		`json:"room_name"`
	Created		time.Time	`json:"created_at"`
	Updated 		time.Time	`json:"updated_at"`
}
// Restriction is the restriction model
type Restriction struct {
	ID 					int			`json:"id"`
	RestrictionName 	string		`json:"restriction_name"`
	Created				time.Time	`json:"created_at"`
	Updated 				time.Time	`json:"updated_at"`
}
// Reservations is the reservation model
type Reservation struct {
	ID 					int			`json:"id"`
	FirstName 			string		`json:"first_name"`
	LastName  			string		`json:"last_name"`
	Email		 			string		`json:"email"`
	Phone					string		`json:"phone"`
	StartDate			time.Time	`json:"start_date"`
	EndDate				time.Time	`json:"end_date"`
	RoomID				int			`json:"room_id"`
	Created				time.Time	`json:"created_at"`
	Updated 				time.Time	`json:"updated_at"`
	Room 					Room			`json:"room"`
	Processed			int			`json:"processed"`
}

// RoomRestrictions is the roomRestrictions model
type RoomRestriction struct {
	ID 					int			`json:"id"`
	StartDate			time.Time	`json:"start_date"`
	EndDate				time.Time	`json:"end_date"`
	RoomID				int			`json:"room_id"`
	ReservationID		int			`json:"reservation_id"`
	RestrictionID		int			`json:"restriction_id"`
	Created				time.Time	`json:"created_at"`
	Updated 				time.Time	`json:"updated_at"`
	Room 					Room			`json:"room"`
	Reservation			Reservation	`json:"reservation"`
	Restriction			Restriction	`json:"restriction"`
}

// MailData holds an email message
type MailData struct {
	To 		string	`json:"to"`
	From 		string	`json:"from"`
	Subject 	string	`json:"subject"`
	Content	string	`json:"content"`
	Template string	`json:"template"`
}
