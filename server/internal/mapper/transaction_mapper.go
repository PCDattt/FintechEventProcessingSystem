package mapper

import (
	"github.com/PCDattt/FintechEventProcessingSystem/server/internal/db"
	"github.com/PCDattt/FintechEventProcessingSystem/shared/enum"
	"github.com/PCDattt/FintechEventProcessingSystem/shared/model"
	"github.com/PCDattt/FintechEventProcessingSystem/shared/proto"
)

func TransactionProtoToModel(tx *proto.TransactionRequest) model.Transaction {
	var toID *int
	if tx.ToAccountId != nil {
		val := int(tx.ToAccountId.Value)
		toID = &val
	}

	return model.Transaction{
		Type: enum.TransactionType(tx.Type),
		Amount: int(tx.Amount),
		AccountId: int(tx.AccountId),
		ToAccountId: toID,
	}
}

func TransactionModelToCreateParams(tx model.Transaction) db.CreateTransactionParams {
	return db.CreateTransactionParams{
		Type: int32(tx.Type),
		Status: int32(tx.Status),
		Amount: int32(tx.Amount),
		Message: tx.Message,
		AccountID: int32(tx.AccountId),
		ToAccountID: tx.ToAccountId,
	}
}

func DBTransactionToModel(tx db.Transaction) model.Transaction {
	return model.Transaction{
		Id: int(tx.ID),
		Type: enum.TransactionType(tx.Type),
		Status: enum.TransactionStatus(tx.Status),
		Amount: int(tx.Amount),
		Message: tx.Message,
		AccountId: int(tx.AccountID),
		ToAccountId: tx.ToAccountID,
		CreatedDate: tx.CreatedDate,
	}
}