CREATE TABLE IF NOT EXISTS tickets (
                                       id          SERIAL PRIMARY KEY,
                                       session_id  INT         NOT NULL REFERENCES sessions(id) ON DELETE CASCADE,
    row         INT         NOT NULL,
    seat        INT         NOT NULL,
    user_id     TEXT        NOT NULL,
    email       TEXT        NOT NULL,
    is_paid     BOOLEAN     NOT NULL DEFAULT FALSE,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
    CONSTRAINT tickets_unique_seat UNIQUE (session_id, row, seat)
    );