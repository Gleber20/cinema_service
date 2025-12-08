package usecase

import (
	"cinema_service/internal/domain"
	"cinema_service/internal/port/driven"
	"cinema_service/internal/port/usecase"
	"errors"
)

type sessionService struct {
	repo driven.SessionRepository
}

func NewSessionService(r driven.SessionRepository) usecase.SessionUseCase {
	return &sessionService{repo: r}
}

func (s *sessionService) ListSessionsByMovie(movieID int) ([]domain.Session, error) {
	return s.repo.GetByMovie(movieID)
}

func (s *sessionService) GetSession(id int) (*domain.Session, error) {
	session, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if session == nil {
		return nil, errors.New("session not found")
	}
	return session, nil
}
