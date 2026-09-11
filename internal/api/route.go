package api

import "github.com/gofiber/fiber/v2"

func (s *Server) Route(route fiber.Router) {
	route.Get("/ready", s.Ready)
	route.Get("/trip", s.GetTrip)
	route.Post("/trip/create", s.CreateTrip)
}
