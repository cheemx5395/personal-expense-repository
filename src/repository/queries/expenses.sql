-- name: CreateExpense :one
INSERT INTO expenses(
    title,
    amount,
    description
) VALUES (
    ?, ?, ?
) RETURNING *;

-- name: GetExpenses :many
SELECT * FROM expenses;

-- name: DeleteExpenseByID :exec
DELETE FROM expenses 
WHERE id = ?;

-- name: UpdateExpenseByID :one
UPDATE expenses
SET title = ?, amount = ?, description = ?
WHERE id = ?
RETURNING *;

-- name: GetExpenseByID :one
SELECT * FROM expenses WHERE id = ?;