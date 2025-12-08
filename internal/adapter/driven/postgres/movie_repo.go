package postgres

import (
	"cinema_service/internal/domain"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type MovieRepo struct {
	db *sqlx.DB
}

func NewMovieRepo(db *sqlx.DB) *MovieRepo {
	return &MovieRepo{db: db}
}

func (mr *MovieRepo) GetAll() ([]domain.Movie, error) {
	var movies []domain.Movie

	err := mr.db.Select(&movies, `
		SELECT id, title, description, duration_min
		FROM movies
		ORDER BY id
	`)
	if err != nil {
		return nil, err
	}

	return movies, nil
}

func (mr *MovieRepo) GetByID(id int) (*domain.Movie, error) {
	var m domain.Movie

	err := mr.db.Get(&m, `
		SELECT id, title, description, duration_min
		FROM movies
		WHERE id = $1
	`, id)
	if err != nil {
		// sqlx просто оборачивает sql.ErrNoRows
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &m, nil
}
