package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"job4j.ru/share_trip/internal/domain"
)

type CreateTripRequest struct {
	ClientID       uuid.UUID
	FromPoint      string
	ToPoint        string
	DepartureTime  time.Time
	AvailableSeats int
}

type CreateTripResponse struct {
	ID             uuid.UUID
	ClientID       uuid.UUID
	FromPoint      string
	ToPoint        string
	DepartureTime  time.Time
	AvailableSeats int
	Status         string
	CreatedAt      time.Time
}

func (s *Service) CreateTrip(
	ctx context.Context,
	request CreateTripRequest,
) (CreateTripResponse, error) {
	domainReq := domain.CreateTripRequest(request)

	createdTrip, err := s.Domain.CreateTrip(ctx, domainReq)
	if err != nil {
		return CreateTripResponse{}, err
	}

	return CreateTripResponse(createdTrip), nil
}
