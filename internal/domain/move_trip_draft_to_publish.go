package domain

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"job4j.ru/share_trip/internal/repository"
)

type MoveTripDraftToPublishModelRequest struct {
	ID       uuid.UUID
	ClientID uuid.UUID
}

type MoveTripDraftToPublishModelResponse struct {
	ID uuid.UUID
}

func (d *Domain) MoveTripDraftToPublish(
	ctx context.Context,
	tx pgx.Tx,
	req MoveTripDraftToPublishModelRequest,
) (*MoveTripDraftToPublishModelResponse, error) {
	trip, err := d.Repository.GetForUpdateByID(ctx, tx, req.ID)
	if err != nil {
		return nil, fmt.Errorf("tripRepository.GetForUpdateByID: %w", err)
	}

	if trip.ClientID != req.ClientID {
		return nil, fmt.Errorf("forbidden: client %s is not driver of trip %s", req.ClientID, req.ID)
	}

	if trip.Status == TripStatusPublished {
		return &MoveTripDraftToPublishModelResponse{
			ID: trip.ID,
		}, nil
	}

	if trip.Status != TripStatusDraft {
		return nil, fmt.Errorf("invalid trip status: expected %s, got %s", TripStatusDraft, trip.Status)
	}

	updatedTrip, err := d.Repository.Update(ctx, tx, trip)
	if err != nil {
		return nil, fmt.Errorf("tripRepository.Update: %w", err)
	}

	draftStatus := string(TripStatusDraft)
	tripHistoryID := uuid.New()

	tripHistory := repository.TripHistory{
		ID:         tripHistoryID,
		TripID:     req.ID,
		FromStatus: &draftStatus,
		ToStatus:   TripStatusPublished,
		CreatedAt:  time.Now(),
	}

	err = d.Repository.CreateTripHistory(ctx, tx, tripHistory)
	if err != nil {
		return nil, fmt.Errorf("tripRepository.Update.CreateTripHistory: %w", err)
	}

	return &MoveTripDraftToPublishModelResponse{
		ID: updatedTrip.ID,
	}, nil
}
