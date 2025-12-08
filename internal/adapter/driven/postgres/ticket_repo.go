package postgres

import (
	"cinema_service/internal/domain"

	"github.com/jmoiron/sqlx"
)

type TicketRepo struct {
	db *sqlx.DB
}

func NewTicketRepo(db *sqlx.DB) *TicketRepo {
	return &TicketRepo{db: db}
}

func (t *TicketRepo) IsSeatBusy(sessionID, row, seat int) (bool, error) {
	var exists bool

	err := t.db.Get(&exists, `
		SELECT EXISTS(
			SELECT 1 FROM tickets
			WHERE session_id = $1 AND row = $2 AND seat = $3
		)
	`, sessionID, row, seat)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (t *TicketRepo) Create(ti domain.Ticket) (int, error) {
	var id int

	err := t.db.QueryRowx(`
		INSERT INTO tickets (session_id, row, seat, user_id, email, is_paid)
		VALUES ($1, $2, $3, $4, $5 ,$6)
		RETURNING id
	`, ti.SessionID, ti.Row, ti.Seat, ti.UserID, ti.Email, ti.IsPaid).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (t *TicketRepo) ListByUser(userID string) ([]domain.Ticket, error) {
	var tickets []domain.Ticket

	err := t.db.Select(&tickets, `
		SELECT id, session_id, row, seat, user_id, email, is_paid, created_at
		FROM tickets
		WHERE user_id = $1
		ORDER BY created_at DESC
	`, userID)
	if err != nil {
		return nil, err
	}

	return tickets, nil
}
