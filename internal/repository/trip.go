package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type Trip struct {
	ID             uuid.UUID
	ClientID       uuid.UUID
	FromPoint      string
	ToPoint        string
	DepartureTime  time.Time
	AvailableSeats int
	Status         string
	CreatedAt      time.Time
}

func (r *RepoPg) GetTrip(ctx context.Context, id uuid.UUID) (Trip, error) {
	var trip Trip
	err := r.pool.QueryRow(
		ctx,
		`SELECT id, client_id, from_point, to_point, departure_time, seats, status, created_at
		 FROM trips WHERE id = $1`,
		id,
	).Scan(
		&trip.ID,
		&trip.ClientID,
		&trip.FromPoint,
		&trip.ToPoint,
		&trip.DepartureTime,
		&trip.AvailableSeats,
		&trip.Status,
		&trip.CreatedAt,
	)
	if err != nil {
		return Trip{}, err
	}
	return trip, nil
}

func (r *RepoPg) CreateTrip(
	ctx context.Context,
	tx pgx.Tx,
	trip Trip,
) (Trip, error) {
	var result Trip
	err := tx.QueryRow(
		ctx,
		`INSERT INTO trips (id, client_id, from_point, to_point, departure_time, seats, status, created_at)
		 VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		 RETURNING id, client_id, from_point, to_point, departure_time, seats, status, created_at`,
		trip.ID,
		trip.ClientID,
		trip.FromPoint,
		trip.ToPoint,
		trip.DepartureTime,
		trip.AvailableSeats,
		trip.Status,
		trip.CreatedAt,
	).Scan(
		&result.ID,
		&result.ClientID,
		&result.FromPoint,
		&result.ToPoint,
		&result.DepartureTime,
		&result.AvailableSeats,
		&result.Status,
		&result.CreatedAt,
	)
	return result, err
}

func (r *RepoPg) GetForUpdateByID(
	ctx context.Context,
	tx pgx.Tx,
	id uuid.UUID,
) (Trip, error) {
	var trip Trip
	err := tx.QueryRow(
		ctx,
		`SELECT id, client_id, from_point, to_point, departure_time, seats, status, created_at
        		 FROM trips WHERE id = $1
        		 FOR UPDATE`,
		id,
	).Scan(
		&trip.ID,
		&trip.ClientID,
		&trip.FromPoint,
		&trip.ToPoint,
		&trip.DepartureTime,
		&trip.AvailableSeats,
		&trip.Status,
		&trip.CreatedAt,
	)
	if err != nil {
		return Trip{}, err
	}
	return trip, err
}

func (r *RepoPg) Update(
	ctx context.Context,
	tx pgx.Tx,
	trip Trip,
) (Trip, error) {
	var result Trip
	err := tx.QueryRow(
		ctx,
		`UPDATE trips
		 SET status = $2
		 WHERE id = $1
		 RETURNING id, client_id, from_point, to_point, departure_time, seats, status, created_at`,
		trip.ID,
		trip.Status,
	).Scan(
		&result.ID,
		&result.ClientID,
		&result.FromPoint,
		&result.ToPoint,
		&result.DepartureTime,
		&result.AvailableSeats,
		&result.Status,
		&result.CreatedAt,
	)
	return result, err
}
