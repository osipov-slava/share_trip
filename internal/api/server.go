package api

import (
	"job4j.ru/share_trip/internal/repository"
)

type Server struct {
	Repository *repository.RepoPg
}

func NewServer(repo *repository.RepoPg) *Server {
	return &Server{Repository: repo}
}
