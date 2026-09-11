package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"job4j.ru/share_trip/internal/domain"
	"job4j.ru/share_trip/internal/repository"
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
	tripID := uuid.New()
	domainReq := domain.CreateTripRequest{
		ClientID:       request.ClientID,
		FromPoint:      request.FromPoint,
		ToPoint:        request.ToPoint,
		DepartureTime:  request.DepartureTime,
		AvailableSeats: request.AvailableSeats,
	}

	domainRes, err := s.Domain.CreateTrip(ctx, domainReq)
	if err != nil {
		return CreateTripResponse{}, err
	}

	trip := repository.Trip{
		ID:             tripID,
		ClientID:       domainRes.ClientID,
		FromPoint:      domainRes.FromPoint,
		ToPoint:        domainRes.ToPoint,
		DepartureTime:  domainRes.DepartureTime,
		AvailableSeats: domainRes.AvailableSeats,
		Status:         "draft",
		CreatedAt:      time.Now(),
	}

	createdTrip, err := s.Repository.CreateTrip(ctx, trip)
	if err != nil {
		return CreateTripResponse{}, err
	}

	tripHistoryID := uuid.New()
	tripHistory := repository.TripHistory{
		ID:         tripHistoryID,
		TripID:     tripID,
		FromStatus: nil,
		ToStatus:   "draft",
		CreatedAt:  time.Now(),
	}

	err = s.Repository.CreateTripHistory(ctx, tripHistory)
	if err != nil {
		return CreateTripResponse{}, err
	}

	return CreateTripResponse{
		ID:             createdTrip.ID,
		ClientID:       createdTrip.ClientID,
		FromPoint:      createdTrip.FromPoint,
		ToPoint:        createdTrip.ToPoint,
		DepartureTime:  createdTrip.DepartureTime,
		AvailableSeats: createdTrip.AvailableSeats,
		Status:         createdTrip.Status,
		CreatedAt:      createdTrip.CreatedAt,
	}, nil
}
