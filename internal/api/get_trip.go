package api

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"job4j.ru/share_trip/internal/service"
)

type GetTripRequest struct {
	TripID uuid.UUID `json:"id"`
}

type GetTripResponse struct {
	ID             uuid.UUID `json:"id"`
	ClientID       uuid.UUID `json:"client_id"`
	FromPoint      string    `json:"from_point"`
	ToPoint        string    `json:"to_point"`
	DepartureTime  time.Time `json:"departure_time"`
	AvailableSeats int       `json:"available_seats"`
	Status         string    `json:"status"`
	CreatedAt      time.Time `json:"created_at"`
}

func (s *Server) GetTrip(c *fiber.Ctx) error {
	tripID := c.Query("id")
	if tripID == "" {
		return fiber.NewError(fiber.StatusBadRequest, "id is required")
	}
	tripUUID, err := uuid.Parse(tripID)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid id")
	}

	res, err := s.Service.GetTrip(c.Context(),
		service.GetTripRequest{
			TripID: tripUUID,
		})
	if err != nil {
		// TODO: add error mapping
		return fiber.NewError(fiber.StatusInternalServerError, "internal server error")
	}

	return c.Status(fiber.StatusOK).JSON(GetTripResponse{
		ID:             res.ID,
		ClientID:       res.ClientID,
		FromPoint:      res.FromPoint,
		ToPoint:        res.ToPoint,
		DepartureTime:  res.DepartureTime,
		AvailableSeats: res.AvailableSeats,
		Status:         res.Status,
		CreatedAt:      res.CreatedAt,
	})
}
