package domain

import "time"

type Session struct {
	ID          int       `db:"id"`
	MovieID     int       `db:"movie_id"`
	StartTime   time.Time `db:"start_time"`
	Price       float64   `db:"price"`
	Rows        int       `db:"rows"`
	SeatsPerRow int       `db:"seats_per_row"`
}
