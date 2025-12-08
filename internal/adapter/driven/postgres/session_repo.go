package postgres

import (
	"cinema_service/internal/domain"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type SessionRepo struct {
	db *sqlx.DB
}

func NewSessionRepo(db *sqlx.DB) *SessionRepo {
	return &SessionRepo{db: db}
}

func (s *SessionRepo) GetByMovie(movieID int) ([]domain.Session, error) {
	var sessions []domain.Session

	err := s.db.Select(&sessions, `
		SELECT id, movie_id, start_time, price, rows, seats_per_row
		FROM sessions
		WHERE movie_id = $1
		ORDER BY start_time
	`, movieID)
	if err != nil {
		return nil, err
	}

	return sessions, nil
}

func (s *SessionRepo) GetByID(id int) (*domain.Session, error) {
	var ses domain.Session

	err := s.db.Get(&ses, `
		SELECT id, movie_id, start_time, price, rows, seats_per_row
		FROM sessions
		WHERE id = $1
	`, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &ses, nil
}
