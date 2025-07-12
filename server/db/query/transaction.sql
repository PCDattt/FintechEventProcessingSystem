-- name: CreateTransaction :one
INSERT INTO transactions (type, status, amount, message, account_id, to_account_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, type, status, amount, message, account_id, to_account_id, created_date;