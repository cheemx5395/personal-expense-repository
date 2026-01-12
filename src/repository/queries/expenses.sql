-- name: CreateExpense :one
INSERT INTO expenses(
    remark,
    amount
) VALUES (
    ?, ?
) RETURNING *;

-- name: GetExpenses :many
SELECT * FROM expenses;

-- name: DeleteExpenseByID :exec
DELETE FROM expenses 
WHERE id = ?;