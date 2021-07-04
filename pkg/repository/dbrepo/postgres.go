package dbrepo

import (
	// built in Golang packages
	"context"
	"errors"
	"time"
	// External packages/dependencies
	"golang.org/x/crypto/bcrypt"
	// My own packages
	"github.com/darkside1809/bookings/pkg/models"
)

// TODO
func(p *postgresDBRepo) GetAllUsers() ([]models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()

	var users []models.User
	query := `
		SELECT id, first_name, last_name, email, access_level, created_at, updated_at FROM users
		ORDER BY created_at ASC`

	rows, err := p.DB.QueryContext(ctx, query)
	if err != nil {
		return users, err
	}
	defer rows.Close()
	
	for rows.Next() {
		var u models.User
		err := rows.Scan(
			&u.ID,
			&u.FirstName,
			&u.LastName,
			&u.Email,
			&u.AccessLevel,
			&u.Created,
			&u.Updated,
		)
		if err != nil {
			return users, err
		}
		users = append(users, u)
	}

	if err = rows.Err(); err != nil {
		return users, err
	}

	return users, nil
}

// DeleteUserByID deletes one reservation by id
func (p *postgresDBRepo) DeleteUserByID(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()

	query := "DELETE FROM users WHERE id = $1"

	_, err := p.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

// GetUserByID returns a user by id
func (p *postgresDBRepo) GetUserByID(id int) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()

	query := `SELECT id, first_name, last_name, email, password, access_level, created_at, updated_at
					FROM users WHERE id = $1`

	row := p.DB.QueryRowContext(ctx, query, id)

	var u models.User
	err := row.Scan(
					&u.ID, 
					&u.FirstName, 
					&u.LastName, 
					&u.Email, 
					&u.Password, 
					&u.AccessLevel, 
					&u.Created, 
					&u.Updated,
	)
	if err != nil {
		return u, err
	}

	return u, nil
}

// UpdateUser updates a user in the database by provided data
func (p *postgresDBRepo) UpdateUser(u models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()

	query := `UPDATE users SET first_name = $1, last_name = $2, email = $3, access_level = $4, updated_at = $5
					WHERE id = $6`
	_, err := p.DB.ExecContext(ctx, query, 
					u.FirstName,
					u.LastName,
					u.Email,
					u.AccessLevel,
					time.Now(),
					u.ID,
	)
	if err != nil {
		return err
	}

	return nil
}

// RegisterUser a user in the system
func (p *postgresDBRepo) RegisterUser(u models.User) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()

	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	var id int
	query := `INSERT INTO users (first_name, last_name, email, password, created_at, updated_at)
					VALUES($1, $2, $3, $4, $5, $6) RETURNING id`

	err = p.DB.QueryRowContext(ctx, query,
					u.FirstName,
					u.LastName,
					u.Email,
					hash,
					time.Now(),
					time.Now(),
	).Scan(&id)
	
	if err != nil {
		return 0, err
	}

	return id, nil
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
				res.RoomID,
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
				res.RoomID,
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

// SearchAvailabilityByDatesByRoomID return true if availability exists for roomID, 
// and false if no abalability exists
func(p *postgresDBRepo) SearchAvailabilityByDatesByRoomID(start time.Time, end time.Time, roomID int) (bool, error){
	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()

	var numRows int

	query := `SELECT COUNT(id) FROM room_restrictions
					WHERE room_id = $1 AND $2 < end_date AND $3 > start_date`

	row := p.DB.QueryRowContext(ctx, query, roomID, start, end)
	err := row.Scan(&numRows)
	if err != nil {
		return false , err
	}
	if numRows == 0 {
		return true, nil
	}

	return false, nil
}

// SearchAvailabilityForAllRooms returns a slice of available rooms, if any, for given date range
func (p *postgresDBRepo) SearchAvailabilityForAllRooms(start time.Time, end time.Time) ([]models.Room, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()
	
	var rooms []models.Room

	query := `SELECT r.id, r.room_name FROM rooms r 
					WHERE r.id NOT IN 
					(SELECT room_id FROM room_restrictions rr WHERE $1 < rr.end_date and $2 > rr.start_date)`

	rows, err := p.DB.QueryContext(ctx, query, start, end)
	if err != nil {
		return rooms , err
	}

	for rows.Next() {
		var room models.Room
		err := rows.Scan(
			&room.ID,
			&room.RoomName,
		)
		if err != nil {
			return rooms, err
		}
		rooms = append(rooms, room)
	}
	if err = rows.Err(); err != nil {
		return rooms, err
	}

	return rooms, nil
}

// GetRoomByID gets a room by id
func (p *postgresDBRepo) GetRoomByID(id int) (models.Room, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()

	var room models.Room
	query := `SELECT id, room_name, created_at, updated_at
					FROM rooms WHERE id = $1`

	row := p.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&room.ID,
		&room.RoomName,
		&room.Created,
		&room.Updated,
	)

	if err != nil {
		return room, err
	}
	return room, nil
}

// Authenticate authenticates a user, if a user exists in database
func (p *postgresDBRepo) Authenticate(email string, password string) (int, string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()

	var id int
	var hash string
	
	row := p.DB.QueryRowContext(ctx, `SELECT id, password FROM users WHERE email = $1`, email)
	err := row.Scan(&id, &hash)
	if err != nil {
		return id, "", err
	}
	
	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return 0, "", errors.New("Incorrect password")
	}
	if err != nil {
		return 0, "", err
	}
	return id, hash , nil
}

// AllReservations returns a slice of all reservations
func (p *postgresDBRepo) GetAllReservations() ([]models.Reservation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()

	var reservations []models.Reservation
	query := `
		SELECT r.id, r.first_name, r.last_name, r.email, r.phone, 
			 r.start_date, r.end_date, r.room_id, r.created_at, r.updated_at, r.processed,
			 rm.id, rm.room_name FROM reservations r
		LEFT JOIN rooms rm ON (r.room_id = rm.id)
		ORDER BY r.start_date ASC`

	rows, err := p.DB.QueryContext(ctx, query)
	if err != nil {
		return reservations, err
	}
	defer rows.Close()
	
	for rows.Next() {
		var i models.Reservation
		err := rows.Scan(
			&i.ID,
			&i.FirstName,
			&i.LastName,
			&i.Email,
			&i.Phone,
			&i.StartDate,
			&i.EndDate,
			&i.RoomID,
			&i.Created,
			&i.Updated,
			&i.Processed,
			&i.Room.ID,
			&i.Room.RoomName,
		)
		if err != nil {
			return reservations, err
		}
		reservations = append(reservations, i)
	}

	if err = rows.Err(); err != nil {
		return reservations, err
	}

	return reservations, nil
}

// AllNewReservations returns a slice of all reservations
func (p *postgresDBRepo) GetAllNewReservations() ([]models.Reservation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()

	var reservations []models.Reservation
	query := `
		SELECT r.id, r.first_name, r.last_name, r.email, r.phone, 
			 r.start_date, r.end_date, r.room_id, r.created_at, r.updated_at,
			 rm.id, rm.room_name FROM reservations r
		LEFT JOIN rooms rm ON (r.room_id = rm.id)
		WHERE processed = 0
		ORDER BY r.start_date ASC`

	rows, err := p.DB.QueryContext(ctx, query)
	if err != nil {
		return reservations, err
	}
	defer rows.Close()
	
	for rows.Next() {
		var i models.Reservation
		err := rows.Scan(
			&i.ID,
			&i.FirstName,
			&i.LastName,
			&i.Email,
			&i.Phone,
			&i.StartDate,
			&i.EndDate,
			&i.RoomID,
			&i.Created,
			&i.Updated,
			&i.Room.ID,
			&i.Room.RoomName,
		)
		if err != nil {
			return reservations, err
		}
		reservations = append(reservations, i)
	}

	if err = rows.Err(); err != nil {
		return reservations, err
	}

	return reservations, nil
}

// GetReservationByID returns one reservation by ID
func (p *postgresDBRepo) GetReservationByID(id int) (models.Reservation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()
	
	var res models.Reservation

	query := `
		SELECT r.id, r.first_name, r.last_name, r.email, r.phone, 
			 r.start_date, r.end_date, r.room_id, r.created_at, r.updated_at, r.processed,
			 rm.id, rm.room_name FROM reservations r
		LEFT JOIN rooms rm ON (r.room_id = rm.id)
		WHERE r.id = $1`

	row := p.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(
				&res.ID,
				&res.FirstName,
				&res.LastName,
				&res.Email,
				&res.Phone,
				&res.StartDate,
				&res.EndDate,
				&res.RoomID,
				&res.Created,
				&res.Updated,
				&res.Processed,
				&res.Room.ID,
				&res.Room.RoomName,
	)
	if err != nil {
		return res, err
	}
	return res, nil
}

// UpdateReservation updates resertion in the database
func (p *postgresDBRepo) UpdateReservation(r models.Reservation) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()

	query := `UPDATE reservations SET first_name = $1, last_name = $2, 
					email = $3, phone = $4, updated_at = $5
				WHERE id = $6`
	
	_, err := p.DB.ExecContext(ctx, query, 
					r.FirstName,
					r.LastName,
					r.Email,
					r.Phone,
					time.Now(),
					r.ID,
	)
	if err != nil {
		return err
	}

	return nil
}

// DeleteReservationByID deletes one reservation by id
func (p *postgresDBRepo) DeleteReservationByID(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()

	query := "DELETE FROM reservations WHERE id = $1"

	_, err := p.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

// UpdateProcessedForReservation updates processed for a reservation by id
func (p *postgresDBRepo) UpdateProcessedForReservation(id int, processed int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()

	query := "UPDATE reservations SET processed = $1 WHERE id = $2"

	_, err := p.DB.ExecContext(ctx, query, processed, id)
	if err != nil {
		return err
	}

	return nil
} 
