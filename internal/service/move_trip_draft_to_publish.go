package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"job4j.ru/share_trip/internal/domain"
)

type MoveTripDraftToPublishModelRequest struct {
	ID       uuid.UUID
	ClientID uuid.UUID
}

type MoveTripDraftToPublishModelResponse struct {
	ID uuid.UUID
}

func (s *Service) MoveTripDraftToPublish(
	ctx context.Context,
	request MoveTripDraftToPublishModelRequest,
) (*MoveTripDraftToPublishModelResponse, error) {
	response, err := tx(ctx, s.Pool, func(tx pgx.Tx) (*MoveTripDraftToPublishModelResponse, error) {
		resp, err := s.Domain.MoveTripDraftToPublish(ctx, tx, domain.MoveTripDraftToPublishModelRequest(request))
		if err != nil {
			return nil, fmt.Errorf("usecase.CreateTrip: %w", err)
		}

		return &MoveTripDraftToPublishModelResponse{
			ID: resp.ID,
		}, nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed in transaction: %w", err)
	}

	return &MoveTripDraftToPublishModelResponse{
		ID: response.ID,
	}, nil
}
