package domain

import "job4j.ru/share_trip/internal/repository"

type Domain struct {
	Repository *repository.RepoPg
}
