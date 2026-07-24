package api

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

func (s *Server) Ready(c *fiber.Ctx) error {
	// Проверяем подключение к БД через пул
	if err := s.Repository.Ping(context.Background()); err != nil {
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
			"status":  "unavailable",
			"message": "Database connection failed",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Service is ready",
		"db":      "Database connection is ready",
	})
}
