CREATE TABLE IF NOT EXISTS movies (
                                      id           SERIAL PRIMARY KEY,
                                      title        TEXT        NOT NULL,
                                      description  TEXT,
                                      duration_min INT         NOT NULL
);