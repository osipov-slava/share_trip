package domain

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"job4j.ru/share_trip/internal/repository"
)

var (
	ErrInvalidAvailableSeats = errors.New("available seats must be >= 1")
	ErrInvalidDepartureTime  = errors.New("departure time must be in the future")
	ErrInvalidPoints         = errors.New("from and to points must be different")
	ErrInvalidClientID       = errors.New("invalid client ID")
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

func (d Domain) CreateTrip(ctx context.Context, req CreateTripRequest) (CreateTripResponse, error) {
	noValidResponse := CreateTripResponse{
		ClientID:       req.ClientID,
		FromPoint:      req.FromPoint,
		ToPoint:        req.ToPoint,
		DepartureTime:  req.DepartureTime,
		AvailableSeats: req.AvailableSeats}
	if req.AvailableSeats < 1 {
		return noValidResponse, ErrInvalidAvailableSeats
	}
	if req.DepartureTime.Before(time.Now()) {
		return noValidResponse, ErrInvalidDepartureTime
	}
	if req.FromPoint == req.ToPoint {
		return noValidResponse, ErrInvalidPoints
	}
	if req.ClientID == uuid.Nil {
		return noValidResponse, ErrInvalidClientID
	}

	tripID := uuid.New()
	trip := repository.Trip{
		ID:             tripID,
		ClientID:       req.ClientID,
		FromPoint:      req.FromPoint,
		ToPoint:        req.ToPoint,
		DepartureTime:  req.DepartureTime,
		AvailableSeats: req.AvailableSeats,
		Status:         "draft",
		CreatedAt:      time.Now(),
	}

	createdTrip, err := d.Repository.CreateTrip(ctx, trip)
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

	err = d.Repository.CreateTripHistory(ctx, tripHistory)
	if err != nil {
		return CreateTripResponse{}, err
	}

	response := CreateTripResponse(createdTrip)
	return response, nil
}
