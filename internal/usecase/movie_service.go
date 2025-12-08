package usecase

import (
	"cinema_service/internal/domain"
	"cinema_service/internal/port/driven"
	"cinema_service/internal/port/usecase"
	"errors"
)

type movieService struct {
	repo driven.MovieRepository
}

func NewMovieService(r driven.MovieRepository) usecase.MovieUseCase {
	return &movieService{repo: r}
}

func (s *movieService) ListMovies() ([]domain.Movie, error) {
	return s.repo.GetAll()
}

func (s *movieService) GetMovie(id int) (*domain.Movie, error) {
	movie, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if movie == nil {
		return nil, errors.New("movie not found")
	}
	return movie, nil
}
