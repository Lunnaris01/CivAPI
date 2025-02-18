-- +goose Up
CREATE TABLE leaders (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL UNIQUE,
    description TEXT NOT NULL,
    country TEXT NOT NULL,
);

-- +goose Down
DROP TABLE leaders;
