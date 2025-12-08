package domain

import "time"

type Ticket struct {
	ID        int       `db:"id"`
	SessionID int       `db:"session_id"`
	Row       int       `db:"row"`
	Seat      int       `db:"seat"`
	UserID    string    `db:"user_id"`
	Email     string    `db:"email"`
	IsPaid    bool      `db:"is_paid"`
	CreatedAt time.Time `db:"created_at"`
}
