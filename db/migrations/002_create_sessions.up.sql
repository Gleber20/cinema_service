CREATE TABLE IF NOT EXISTS sessions (
                                        id             SERIAL PRIMARY KEY,
                                        movie_id       INT         NOT NULL REFERENCES movies(id) ON DELETE CASCADE,
    start_time     TIMESTAMPTZ NOT NULL,
    rows           INT         NOT NULL,
    seats_per_row  INT         NOT NULL,
    created_at     TIMESTAMPTZ NOT NULL DEFAULT now()
    );