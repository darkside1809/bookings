package dbrepo

import (
	"errors"
	"time"

	"github.com/darkside1809/bookings/internal/models"
)

func(p *testDBRepo) AllUsers() bool {
	return true
}
// InsertReservation inserts a reservations into the database
func(p *testDBRepo) InsertReservation(res models.Reservation) (int, error) {
	if res.RoomID == 2 {
		return 0, errors.New("Invalid insertion")
	}
	return 1, nil
}
// InsertRoomRestriction inserts a room restrictions into the database
func(p *testDBRepo) InsertRoomRestriction(res models.RoomRestriction) error {
	if res.RoomID == 1000 {
		return errors.New("invalid insertion")
	}
	return nil
}
// SearchAvailabilityByDatesByRoomID return true if availability exists for roomID, 
// and false if no abalability exists
func(p *testDBRepo) SearchAvailabilityByDatesByRoomID(start time.Time, end time.Time, roomID int) (bool, error){
	return false, nil
}
// SearchAvailabilityForAllRooms returns a slice of available rooms, if any, for given date range
func (p *testDBRepo) SearchAvailabilityForAllRooms(start time.Time, end time.Time) ([]models.Room, error) {
	var rooms []models.Room
	return rooms, nil
}
// GetRoomByID gets a room by id
func (p *testDBRepo) GetRoomByID(id int) (models.Room, error) {
	var room models.Room
	if id > 2 {
		return room, errors.New("Some error")
	}

	return room, nil
}