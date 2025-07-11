-- name: CreateTransaction :one
INSERT INTO transactions (type, status, amount, message, from_Account_Id, to_Account_Id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, status, message;