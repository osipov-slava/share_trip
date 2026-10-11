package api

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"job4j.ru/share_trip/internal/domain"
	"job4j.ru/share_trip/internal/repository"
	"job4j.ru/share_trip/internal/service"
)

type Server struct {
	Repository *repository.RepoPg
	Service    *service.Service
}

func NewServer(repo *repository.RepoPg, pool *pgxpool.Pool) *Server {
	dom := &domain.Domain{
		Repository: repo,
	}
	svc := &service.Service{
		Domain: dom,
		Pool:   pool,
	}
	return &Server{
		Repository: repo,
		Service:    svc,
	}
}
