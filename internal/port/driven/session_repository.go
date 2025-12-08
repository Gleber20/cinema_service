package driven

import "cinema_service/internal/domain"

type SessionRepository interface {
	GetByMovie(movieID int) ([]domain.Session, error)
	GetByID(id int) (*domain.Session, error)
}
