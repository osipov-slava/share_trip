package service

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

func (s *Service) GetTrip(
	ctx context.Context,
	request GetTripRequest,
) (GetTripResponse, error) {
	trip, err := s.Repository.GetTrip(ctx, request.TripID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return GetTripResponse{}, err
		}
		return GetTripResponse{}, err
	}

	return GetTripResponse{
		ID:             trip.ID,
		ClientID:       trip.ClientID,
		FromPoint:      trip.FromPoint,
		ToPoint:        trip.ToPoint,
		DepartureTime:  trip.DepartureTime,
		AvailableSeats: trip.AvailableSeats,
		Status:         trip.Status,
		CreatedAt:      trip.CreatedAt,
	}, nil
}
