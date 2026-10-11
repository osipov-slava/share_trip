package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type TripHistory struct {
	ID         uuid.UUID
	TripID     uuid.UUID
	FromStatus *string
	ToStatus   string
	CreatedAt  time.Time
}

func (r *RepoPg) CreateTripHistory(
	ctx context.Context,
	tx pgx.Tx,
	tripHistory TripHistory,
) error {
	return tx.QueryRow(
		ctx,
		`INSERT INTO trip_history (id, trip_id, from_status, to_status, created_at)
		 VALUES ($1, $2, $3, $4, $5)
		 RETURNING id, trip_id, from_status, to_status, created_at`,
		tripHistory.ID,
		tripHistory.TripID,
		tripHistory.FromStatus,
		tripHistory.ToStatus,
		tripHistory.CreatedAt,
	).Scan(
		&tripHistory.ID,
		&tripHistory.TripID,
		&tripHistory.FromStatus,
		&tripHistory.ToStatus,
		&tripHistory.CreatedAt,
	)
}
