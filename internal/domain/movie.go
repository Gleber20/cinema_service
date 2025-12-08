package domain

type Movie struct {
	ID          int    `db:"id"`
	Title       string `db:"title"`
	Description string `db:"description"`
	DurationMin int    `db:"duration_min"`
}
