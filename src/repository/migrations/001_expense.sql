-- +goose Up
CREATE TABLE expenses(
    id INTEGER PRIMARY KEY,
    remark TEXT,
    amount INTEGER NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE expenses;