package service

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"job4j.ru/share_trip/internal/domain"
)

type Service struct {
	Domain *domain.Domain
	Pool   *pgxpool.Pool
}
