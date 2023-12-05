package dbrepo

import (
	"context"
	"time"

	"github.com/manuel-valles/bookings-app.git/internal/models"
)

func (pr *postgresDBRepo) InsertReservation(res models.Reservation) (int, error) {
	// Make sure that DB transaction doesn't stay open during all the maxDbLifetime
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
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
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
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
