// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: transaction.sql

package db

import (
	"context"
)

const createTransaction = `-- name: CreateTransaction :one
INSERT INTO transactions (type, status, amount, message, account_id, to_account_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, type, status, amount, message, account_id, to_account_id, created_date
`

type CreateTransactionParams struct {
	Type        int32  `json:"type"`
	Status      int32  `json:"status"`
	Amount      int32  `json:"amount"`
	Message     string `json:"message"`
	AccountID   int32  `json:"account_id"`
	ToAccountID *int   `json:"to_account_id"`
}

func (q *Queries) CreateTransaction(ctx context.Context, arg CreateTransactionParams) (Transaction, error) {
	row := q.db.QueryRow(ctx, createTransaction,
		arg.Type,
		arg.Status,
		arg.Amount,
		arg.Message,
		arg.AccountID,
		arg.ToAccountID,
	)
	var i Transaction
	err := row.Scan(
		&i.ID,
		&i.Type,
		&i.Status,
		&i.Amount,
		&i.Message,
		&i.AccountID,
		&i.ToAccountID,
		&i.CreatedDate,
	)
	return i, err
}

const updateTransaction = `-- name: UpdateTransaction :one
UPDATE transactions 
SET status = $2, message = $3
WHERE id = $1
RETURNING id, type, status, amount, message, account_id, to_account_id, created_date
`

type UpdateTransactionParams struct {
	ID      int32  `json:"id"`
	Status  int32  `json:"status"`
	Message string `json:"message"`
}

func (q *Queries) UpdateTransaction(ctx context.Context, arg UpdateTransactionParams) (Transaction, error) {
	row := q.db.QueryRow(ctx, updateTransaction, arg.ID, arg.Status, arg.Message)
	var i Transaction
	err := row.Scan(
		&i.ID,
		&i.Type,
		&i.Status,
		&i.Amount,
		&i.Message,
		&i.AccountID,
		&i.ToAccountID,
		&i.CreatedDate,
	)
	return i, err
}
