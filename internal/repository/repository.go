package repository

import (
	// built in Golang packages
	"time"
	// My own packages
	"github.com/darkside1809/bookings/internal/models"
)
// DatabaseRepo is an interface that must be implemented by postgresDBRepo
type DatabaseRepo interface {
	InsertReservation(res models.Reservation) (int, error)
	InsertRoomRestriction(res models.RoomRestriction) error

	SearchAvailabilityByDatesByRoomID(start time.Time, end time.Time, roomID int) (bool, error)
	SearchAvailabilityForAllRooms(start time.Time, end time.Time) ([]models.Room, error)
	GetRoomByID(id int) (models.Room, error)
	AllRooms() ([]models.Room, error)
	GetRestrictionsForRoomByDate(roomID int, start, end time.Time) ([]models.RoomRestriction, error)
	InsertBlockForRoom(id int, startDate time.Time) error 
	DeleteBlockByID(id int) error

	GetAllUsers() ([]models.User, error) 
	DeleteUserByID(id int) error
	GetUserByID(id int) (models.User, error)
	UpdateUser(u models.User) error
	Authenticate(email string, password string) (int, string, error)
	RegisterUser(u models.User) (int, error)

	GetAllReservations() ([]models.Reservation, error)
	GetAllNewReservations() ([]models.Reservation, error)
	GetReservationByID(id int) (models.Reservation, error)
	UpdateReservation(r models.Reservation) error
	DeleteReservationByID(id int) error
	UpdateProcessedForReservation(id int, processed int) error
}