package dbrepo

import (
	"context"
	"time"

	"github.com/manuel-valles/bookings-app.git/internal/models"
)

const maxQueryLife = 3 * time.Second

func (pr *postgresDBRepo) InsertReservation(res models.Reservation) (int, error) {
	// Make sure that DB transaction doesn't stay open during all the maxDbLifetime
	ctx, cancel := context.WithTimeout(context.Background(), maxQueryLife)
	defer cancel()

	var newID int
	query := `INSERT INTO reservations (first_name, last_name, email, phone, start_date, end_date, room_id, created_at, updated_at) 
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id`

	// Using QueryRowContext() instead of ExecContext() because it does not only insert but also returns a value
	err := pr.DB.QueryRowContext(ctx, query,
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

func (pr *postgresDBRepo) InsertRoomRestriction(rr models.RoomRestriction) error {
	ctx, cancel := context.WithTimeout(context.Background(), maxQueryLife)
	defer cancel()

	query := `INSERT INTO room_restrictions (start_date, end_date, room_id, reservation_id, restriction_id, created_at, updated_at) 
              VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`

	_, err := pr.DB.ExecContext(ctx, query,
		rr.StartDate,
		rr.EndDate,
		rr.RoomID,
		rr.ReservationID,
		rr.RestrictionID,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		return err
	}
	return nil
}

func (pr *postgresDBRepo) SearchAvailabilityByDatesByRoomID(start, end time.Time, roomID int) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), maxQueryLife)
	defer cancel()

	var numRows int

	query := `SELECT COUNT(id)
			  FROM room_restrictions
			  WHERE room_id = $1 AND $2 < end_date AND $3 > start_date;`

	row := pr.DB.QueryRowContext(ctx, query, roomID, start, end)
	err := row.Scan(&numRows)
	if err != nil {
		return false, err
	}

	if numRows == 0 {
		return true, nil
	}
	return false, nil
}

func (pr *postgresDBRepo) SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error) {
	ctx, cancel := context.WithTimeout(context.Background(), maxQueryLife)
	defer cancel()

	var rooms []models.Room

	query := `SELECT r.id, r.room_name
			  FROM rooms r
			  WHERE r.id NOT IN (SELECT room_id FROM room_restrictions rr WHERE $1 < rr.end_date AND $2 > rr.start_date);`

	rows, err := pr.DB.QueryContext(ctx, query, start, end)
	if err != nil {
		return rooms, err
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
