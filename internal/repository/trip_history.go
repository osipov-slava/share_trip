package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
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
	tripHistory TripHistory,
) error {
	return r.pool.QueryRow(
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
