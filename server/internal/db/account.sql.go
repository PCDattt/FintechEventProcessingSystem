// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: account.sql

package db

import (
	"context"
)

const createAccount = `-- name: CreateAccount :one
INSERT INTO accounts (username, password)
VALUES ($1, $2)
RETURNING id, username, password, amount, created_date
`

type CreateAccountParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (q *Queries) CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error) {
	row := q.db.QueryRow(ctx, createAccount, arg.Username, arg.Password)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Amount,
		&i.CreatedDate,
	)
	return i, err
}

const getAccountForUpdate = `-- name: GetAccountForUpdate :one
SELECT id, username, password, amount, created_date FROM accounts
WHERE id = $1
FOR UPDATE
`

func (q *Queries) GetAccountForUpdate(ctx context.Context, id int32) (Account, error) {
	row := q.db.QueryRow(ctx, getAccountForUpdate, id)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Amount,
		&i.CreatedDate,
	)
	return i, err
}

const updateAccountBalance = `-- name: UpdateAccountBalance :one
UPDATE accounts
SET amount = $2
WHERE id = $1
RETURNING id, username, password, amount, created_date
`

type UpdateAccountBalanceParams struct {
	ID     int32 `json:"id"`
	Amount int32 `json:"amount"`
}

func (q *Queries) UpdateAccountBalance(ctx context.Context, arg UpdateAccountBalanceParams) (Account, error) {
	row := q.db.QueryRow(ctx, updateAccountBalance, arg.ID, arg.Amount)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Amount,
		&i.CreatedDate,
	)
	return i, err
}
