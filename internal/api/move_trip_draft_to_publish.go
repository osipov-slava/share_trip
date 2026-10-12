package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"job4j.ru/share_trip/internal/service"
)

type MoveTripDraftToPublishModelRequest struct {
	ID       uuid.UUID `json:"id"`
	ClientID uuid.UUID `json:"client_id"`
}

type MoveTripDraftToPublishModelResponse struct {
	ID uuid.UUID `json:"id"`
}

func (s *Server) MoveTripDraftToPublish(c *fiber.Ctx) error {
	var req MoveTripDraftToPublishModelRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid JSON body")
	}

	res, err := s.Service.MoveTripDraftToPublish(c.Context(),
		service.MoveTripDraftToPublishModelRequest(req))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(MoveTripDraftToPublishModelResponse(*res))
}
