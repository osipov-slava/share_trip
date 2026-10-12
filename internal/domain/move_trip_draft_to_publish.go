package domain

import (
	"context"
	"encoding/json"
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
		return nil, fmt.Errorf("TripRepository.GetForUpdateByID: %w", err)
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

	trip.Status = TripStatusPublished
	updatedTrip, err := d.Repository.Update(ctx, tx, trip)
	if err != nil {
		return nil, fmt.Errorf("TripRepository.Update: %w", err)
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
		return nil, fmt.Errorf("TripHistoryRepository.CreateTripHistory: %w", err)
	}

	payload, err := json.Marshal(map[string]string{
		"task": req.ID.String(),
	})
	if err != nil {
		return nil, fmt.Errorf("domain.MoveTripDraftToPublish json.Marshal: %w", err)
	}

	outboxEvent := repository.OutboxEvent{
		ID:          tripHistoryID,
		EventName:   "trip_published",
		AggregateID: req.ID,
		Payload:     string(payload), // `{"task":"..."}`
		CreatedAt:   time.Now(),
	}

	err = d.Repository.CreateOutboxEvent(ctx, tx, outboxEvent)
	if err != nil {
		return nil, fmt.Errorf("repository.CreateOutboxEvent: %w", err)
	}

	return &MoveTripDraftToPublishModelResponse{
		ID: updatedTrip.ID,
	}, nil
}
