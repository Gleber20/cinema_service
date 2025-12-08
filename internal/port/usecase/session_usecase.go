package usecase

import "cinema_service/internal/domain"

type SessionUseCase interface {
	ListSessionsByMovie(movieID int) ([]domain.Session, error)
	GetSession(id int) (*domain.Session, error)
}
