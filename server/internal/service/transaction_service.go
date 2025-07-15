package service

import (
	"context"
	"fmt"
	"time"

	"github.com/PCDattt/FintechEventProcessingSystem/server/internal/db"
	"github.com/PCDattt/FintechEventProcessingSystem/server/internal/mapper"
	"github.com/PCDattt/FintechEventProcessingSystem/shared/enum"
	"github.com/PCDattt/FintechEventProcessingSystem/shared/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TransactionService interface {
	CreateTransaction(ctx context.Context, tx model.Transaction) (model.Transaction, error)
	ProcessTransaction(ctx context.Context, tx model.Transaction) error
}

type transactionService struct {
	pool *pgxpool.Pool
	q *db.Queries
}

func NewTransactionService(pool *pgxpool.Pool, q *db.Queries) *transactionService {
	return &transactionService{
		pool: pool,
		q: q,
	}
}

func (s *transactionService) CreateTransaction(ctx context.Context, tx model.Transaction) (model.Transaction, error) {
	tx.Status = enum.TransactionStatusProcessing
	tx.Message = "Processing"
	params := mapper.TransactionModelToCreateParams(tx)

	dbTransaction, err := s.q.CreateTransaction(ctx, params)
	if err != nil {
		return model.Transaction{}, err
	}

	model := mapper.DBTransactionToModel(dbTransaction)
	return model, nil
}

func (s *transactionService) ProcessTransaction(ctx context.Context, tx model.Transaction) (model.Transaction, error) {
	time.Sleep(5 * time.Second)
	dbtx, err := s.pool.Begin(ctx)
	if err != nil {
		return model.Transaction{}, err
	}
	defer dbtx.Rollback(ctx)

	qtx := db.New(dbtx)
	account, err := qtx.GetAccountForUpdate(ctx, int32(tx.AccountId))
	if err != nil {
		return model.Transaction{}, err
	}

	newBalance := account.Amount
	if tx.Type == enum.TransactionTypeWithdraw {
		if account.Amount < int32(tx.Amount) {
			return model.Transaction{}, fmt.Errorf("insufficient balance")
		}
		newBalance -= int32(tx.Amount)
	}

	if tx.Type == enum.TransactionTypeDeposit {
		newBalance += int32(tx.Amount)
	}

	if tx.Type == enum.TransactionTypePayment {
		if account.Amount < int32(tx.Amount) {
			return model.Transaction{}, fmt.Errorf("insufficient balance")
		}
		newBalance -= int32(tx.Amount)
		toAccount, err := qtx.GetAccountForUpdate(ctx, int32(*tx.ToAccountId))
		if err != nil {
			return model.Transaction{}, err
		}
		toAccountNewBalance := toAccount.Amount + int32(tx.Amount)
		toAccount, err = qtx.UpdateAccountBalance(ctx, db.UpdateAccountBalanceParams{
			ID: int32(*tx.ToAccountId),
			Amount: toAccountNewBalance,
		})
		if err != nil {
			return model.Transaction{}, err
		}
	}

	account, err = qtx.UpdateAccountBalance(ctx, db.UpdateAccountBalanceParams{
		ID: int32(tx.AccountId),
		Amount: newBalance,
	})
	if err != nil {
		return model.Transaction{}, err
	}

	dbTransaction, err := qtx.UpdateTransaction(ctx, db.UpdateTransactionParams{
		ID: int32(tx.Id),
		Status: int32(enum.TransactionStatusSuccess),
		Message: "Processed",
	})
	if err != nil {
		return model.Transaction{}, err
	}

	model := mapper.DBTransactionToModel(dbTransaction)
	return model, dbtx.Commit(ctx)
}