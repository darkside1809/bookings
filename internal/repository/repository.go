package repository

import (
	"time"
	"github.com/darkside1809/bookings/internal/models"
)
// DatabaseRepo is an interface that must be implemented by postgresDBRepo structure type
type DatabaseRepo interface {
	AllUsers() bool
	InsertReservation(res models.Reservation) (int, error)
	InsertRoomRestriction(res models.RoomRestriction) error
	SearchAvailabilityByDatesByRoomID(start time.Time, end time.Time, roomID int) (bool, error)
	SearchAvailabilityForAllRooms(start time.Time, end time.Time) ([]models.Room, error)
	GetRoomByID(id int) (models.Room, error)
}