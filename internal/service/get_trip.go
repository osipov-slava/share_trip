package service

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"job4j.ru/share_trip/internal/domain"
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
	trip, err := s.Domain.GetTrip(ctx, domain.GetTripRequest(request))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return GetTripResponse{}, err
		}
		return GetTripResponse{}, err
	}

	return GetTripResponse(trip), nil
}
