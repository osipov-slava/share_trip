package service

import (
	"job4j.ru/share_trip/internal/domain"
	"job4j.ru/share_trip/internal/repository"
)

type Service struct {
	Domain     *domain.Domain
	Repository *repository.RepoPg
}
