-- +goose Up
CREATE TABLE expenses(
    id INTEGER PRIMARY KEY,
    title TEXT,
    amount INTEGER NOT NULL,
    description TEXT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE expenses;