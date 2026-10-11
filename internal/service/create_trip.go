package service

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
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
	response, err := tx(ctx, s.Pool, func(tx pgx.Tx) (*CreateTripResponse, error) {
		resp, err := s.Domain.CreateTrip(ctx, tx, domain.CreateTripRequest(request))
		if err != nil {
			return nil, fmt.Errorf("usecase.CreateTrip: %w", err)
		}

		return &CreateTripResponse{
			ID:             resp.ID,
			ClientID:       resp.ClientID,
			FromPoint:      resp.FromPoint,
			ToPoint:        resp.ToPoint,
			DepartureTime:  resp.DepartureTime,
			AvailableSeats: resp.AvailableSeats,
			Status:         resp.Status,
			CreatedAt:      resp.CreatedAt,
		}, nil
	})

	if err != nil {
		return CreateTripResponse{}, fmt.Errorf("failed in transaction: %w", err)
	}

	return *response, nil
}
