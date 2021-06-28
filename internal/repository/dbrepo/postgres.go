package dbrepo

import (
	"context"
	"time"

	"github.com/darkside1809/bookings/internal/models"
)

func(p *postgresDBRepo) AllUsers() bool {
	return true
}
// InsertReservation inserts a reservations into the database
func(p *postgresDBRepo) InsertReservation(res models.Reservation) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()

	var newID int

	stmt := `INSERT INTO reservations (first_name, last_name, 
					email, phone, 
					start_date, end_date, 
					room_id, created_at, updated_at)
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id`

	err := p.DB.QueryRowContext(ctx, stmt, 
				res.FirstName,
				res.LastName,
				res.Email,
				res.Phone,
				res.StartDate,
				res.EndDate,
				res.RoomId,
				time.Now(),
				time.Now(),
	).Scan(&newID)

	if err != nil {
		return 0, err
	}

	return newID, nil
}
// InsertRoomRestriction inserts a room restrictions into the database
func(p *postgresDBRepo) InsertRoomRestriction(res models.RoomRestriction) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()

	stmt := `INSERT INTO room_restrictions (start_date, end_date, 
					room_id, reservation_id, restriction_id, created_at, updated_at)
				VALUES ($1, $2, $3, $4, $5, $6, $7)`
	
	_, err := p.DB.ExecContext(ctx, stmt, 
				res.StartDate,
				res.EndDate,
				res.RoomId,
				res.ReservationID,
				res.RestrictionID,
				time.Now(),
				time.Now(),					
	)
	if err != nil {
		return err
	}
	return nil
}