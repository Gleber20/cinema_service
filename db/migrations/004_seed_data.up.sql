-- Фильмы
INSERT INTO movies (title, description, duration_min) VALUES
                                                          ('Матрица', 'Киберпанк, Нео, выбор красной таблетки', 136),
                                                          ('Начало', 'Погружение во сны, ди Каприо, нолановский мозговзрыв', 148);

-- Сеансы для Матрицы
INSERT INTO sessions (movie_id, start_time, rows, seats_per_row) VALUES
                                                                     (1, now() + interval '1 hour', 5, 10),
                                                                     (1, now() + interval '3 hours', 5, 10);

-- Сеансы для Начала
INSERT INTO sessions (movie_id, start_time, rows, seats_per_row) VALUES
                                                                     (2, now() + interval '2 hours', 6, 12),
                                                                     (2, now() + interval '5 hours', 6, 12);