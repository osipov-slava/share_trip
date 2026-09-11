package api

import (
	"job4j.ru/share_trip/internal/domain"
	"job4j.ru/share_trip/internal/repository"
	"job4j.ru/share_trip/internal/service"
)

type Server struct {
	Repository *repository.RepoPg
	Service    *service.Service
}

func NewServer(repo *repository.RepoPg) *Server {
	svc := &service.Service{
		Domain:     &domain.Domain{},
		Repository: repo,
	}
	return &Server{
		Repository: repo,
		Service:    svc,
	}
}
