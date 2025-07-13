-- name: CreateAccount :one
INSERT INTO accounts (username, password)
VALUES ($1, $2)
RETURNING *;

-- name: GetAccountForUpdate :one
SELECT * FROM accounts
WHERE id = $1
FOR UPDATE;

-- name: UpdateAccountBalance :one
UPDATE accounts
SET amount = $2
WHERE id = $1
RETURNING *;