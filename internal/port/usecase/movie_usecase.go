package usecase

import "cinema_service/internal/domain"

type MovieUseCase interface {
	ListMovies() ([]domain.Movie, error)
	GetMovie(id int) (*domain.Movie, error)
}
