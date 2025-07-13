package service

import (
	"context"
	"log"

	"github.com/PCDattt/FintechEventProcessingSystem/server/internal/db"
	"github.com/PCDattt/FintechEventProcessingSystem/server/internal/mapper"
	"github.com/PCDattt/FintechEventProcessingSystem/shared/enum"
	"github.com/PCDattt/FintechEventProcessingSystem/shared/model"
)

type TransactionService interface {
	CreateTransaction(ctx context.Context, tx model.Transaction) (model.Transaction, error)
	ProcessTransaction(ctx context.Context, tx model.Transaction) (model.Transaction, error)
}

type transactionService struct {
	q *db.Queries
}

func NewTransactionService(q *db.Queries) *transactionService {
	return &transactionService{q: q}
}

func (s *transactionService) CreateTransaction(ctx context.Context, tx model.Transaction) (model.Transaction, error) {
	tx.Status = enum.TransactionStatusProcessing
	tx.Message = "Processing"
	params := mapper.TransactionModelToCreateParams(tx)

	dbTransaction, err := s.q.CreateTransaction(ctx, params)
	if err != nil {
		log.Printf("Cannot create transaction in database: %v", err)
		return model.Transaction{}, err
	}

	model := mapper.DBTransactionToModel(dbTransaction)
	return model, nil
}

func (s *transactionService) ProcessTransaction(ctx context.Context, tx model.Transaction) (model.Transaction, error) {
	return model.Transaction{}, nil
}