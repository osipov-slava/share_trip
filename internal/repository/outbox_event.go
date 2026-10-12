package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type OutboxEvent struct {
	ID          uuid.UUID
	EventName   string
	AggregateID uuid.UUID
	Payload     string
	CreatedAt   time.Time
}

func (r *RepoPg) CreateOutboxEvent(
	ctx context.Context,
	tx pgx.Tx,
	outboxEvent OutboxEvent,
) error {
	return tx.QueryRow(
		ctx,
		`INSERT INTO outbox_event (id,event_name,aggregate_id,payload,created_at)
		 VALUES ($1, $2, $3, $4, $5)
		 RETURNING id,event_name,aggregate_id,payload,created_at`,
		outboxEvent.ID,
		outboxEvent.EventName,
		outboxEvent.AggregateID,
		outboxEvent.Payload,
		outboxEvent.CreatedAt,
	).Scan(
		&outboxEvent.ID,
		&outboxEvent.EventName,
		&outboxEvent.AggregateID,
		&outboxEvent.Payload,
		&outboxEvent.CreatedAt,
	)
}
