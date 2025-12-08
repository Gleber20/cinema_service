package driven

import "cinema_service/internal/domain"

type MovieRepository interface {
	GetAll() ([]domain.Movie, error)
	GetByID(id int) (*domain.Movie, error)
}
