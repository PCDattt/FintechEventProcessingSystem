-- name: CreateAccount :one
INSERT INTO accounts (username, password)
VALUES ($1, $2)
RETURNING id, username, password, amount, created_date;