package dbrepo

import (
	"errors"
	"time"

	"github.com/darkside1809/bookings/pkg/models"
)

func(p *testDBRepo) GetAllUsers() (users []models.User, err error) {
	return users, err
}
func(p *testDBRepo) DeleteUserByID(id int) error {
	return nil
}

func(p *testDBRepo) InsertReservation(res models.Reservation) (int, error) {
	if res.RoomID == 2 {
		return 0, errors.New("Invalid insertion")
	}
	return 1, nil
}

func(p *testDBRepo) InsertRoomRestriction(res models.RoomRestriction) error {
	if res.RoomID == 1000 {
		return errors.New("invalid insertion")
	}
	return nil
}

func(p *testDBRepo) SearchAvailabilityByDatesByRoomID(start time.Time, end time.Time, roomID int) (bool, error){
	return false, nil
}

func (p *testDBRepo) SearchAvailabilityForAllRooms(start time.Time, end time.Time) ([]models.Room, error) {
	var rooms []models.Room
	return rooms, nil
}
func (p *testDBRepo) GetRoomByID(id int) (models.Room, error) {
	var room models.Room
	if id > 2 {
		return room, errors.New("Some error")
	}

	return room, nil
}
func (p *testDBRepo) GetUserByID(id int) (models.User, error) {
	var u models.User
	return u, nil
}
func (p *testDBRepo) UpdateUser(u models.User) error {
	return nil
}

func (p *testDBRepo) Authenticate(email string, password string) (int, string, error) {
	return 1, "", nil
}
func (p *testDBRepo) GetAllReservations() ([]models.Reservation, error) {
	var reservations []models.Reservation

	return reservations, nil
}
func (p *testDBRepo) GetAllNewReservations() ([]models.Reservation, error) {
	var reservations []models.Reservation

	return reservations, nil
}
func (p *testDBRepo) GetReservationByID(id int) (models.Reservation, error) {
	var res models.Reservation

	return res, nil
}
func (p *testDBRepo) UpdateReservation(r models.Reservation) error {
	return nil
}

func (p *testDBRepo) DeleteReservationByID(id int) error {
	return nil
}
func (p *testDBRepo) RegisterUser(u models.User) (int, error) {
	var id int
	return id, nil
} 

func (p *testDBRepo) UpdateProcessedForReservation(id int, processed int) error {
	return nil
} 
 

