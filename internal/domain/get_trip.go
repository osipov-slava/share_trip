package domain

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type GetTripRequest struct {
	TripID uuid.UUID
}

type GetTripResponse struct {
	ID             uuid.UUID
	ClientID       uuid.UUID
	FromPoint      string
	ToPoint        string
	DepartureTime  time.Time
	AvailableSeats int
	Status         string
	CreatedAt      time.Time
}

func (d *Domain) GetTrip(
	ctx context.Context,
	req GetTripRequest,
) (GetTripResponse, error) {
	trip, err := d.Repository.GetTrip(ctx, req.TripID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return GetTripResponse{}, err
		}
		return GetTripResponse{}, err
	}

	return GetTripResponse(trip), nil
}
