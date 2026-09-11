package domain

import (
	"time"

	"github.com/google/uuid"
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
