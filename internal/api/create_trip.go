package api

import (
	"errors"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
	"job4j.ru/share_trip/internal/domain"
	"job4j.ru/share_trip/internal/service"
)

type CreateTripRequest struct {
	ClientID       uuid.UUID `json:"client_id"`
	FromPoint      string    `json:"from_point"`
	ToPoint        string    `json:"to_point"`
	DepartureTime  time.Time `json:"departure_time"`
	AvailableSeats int       `json:"available_seats"`
}

type CreateTripResponse struct {
	ID             uuid.UUID `json:"id"`
	ClientID       uuid.UUID `json:"client_id"`
	FromPoint      string    `json:"from_point"`
	ToPoint        string    `json:"to_point"`
	DepartureTime  time.Time `json:"departure_time"`
	AvailableSeats int       `json:"available_seats"`
	Status         string    `json:"status"`
	CreatedAt      time.Time `json:"created_at"`
}

func (s *Server) CreateTrip(c *fiber.Ctx) error {
	var req CreateTripRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid JSON body")
	}
	if req.FromPoint == "" {
		return fiber.NewError(fiber.StatusBadRequest, "from_point is required")
	}
	if req.ToPoint == "" {
		return fiber.NewError(fiber.StatusBadRequest, "to_point is required")
	}

	res, err := s.Service.CreateTrip(c.Context(),
		service.CreateTripRequest(req))
	if err != nil {
		log.Errorw("s.Service.CreateTrip", err)
		if errors.Is(err, domain.ErrInvalidAvailableSeats) ||
			errors.Is(err, domain.ErrInvalidDepartureTime) ||
			errors.Is(err, domain.ErrInvalidPoints) ||
			errors.Is(err, domain.ErrInvalidClientID) {
			msg := err.Error()
			if i := strings.LastIndex(msg, ": "); i != -1 {
				msg = msg[i+2:]
			}
			return fiber.NewError(fiber.StatusBadRequest, msg)
		}
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(CreateTripResponse(res))
}
