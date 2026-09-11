package domain

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
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
	ClientID       uuid.UUID
	FromPoint      string
	ToPoint        string
	DepartureTime  time.Time
	AvailableSeats int
}

func (d Domain) CreateTrip(ctx context.Context, req CreateTripRequest) (CreateTripResponse, error) {
	response := CreateTripResponse(req)
	if req.AvailableSeats < 1 {
		return response, ErrInvalidAvailableSeats
	}
	if req.DepartureTime.Before(time.Now()) {
		return response, ErrInvalidDepartureTime
	}
	if req.FromPoint == req.ToPoint {
		return response, ErrInvalidPoints
	}
	if req.ClientID == uuid.Nil {
		return response, ErrInvalidClientID
	}
	return response, nil
}
