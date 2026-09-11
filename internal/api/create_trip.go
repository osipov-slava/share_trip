package api

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
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
		service.CreateTripRequest{
			ClientID:       req.ClientID,
			FromPoint:      req.FromPoint,
			ToPoint:        req.ToPoint,
			DepartureTime:  req.DepartureTime,
			AvailableSeats: req.AvailableSeats,
		})
	if err != nil {
		//TODO: обработать виды ошибок
		log.Errorw("s.Service.CreateTrip", err)
		return fiber.NewError(fiber.StatusInternalServerError, "internal server error")
	}

	return c.Status(fiber.StatusCreated).JSON(CreateTripResponse{
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
