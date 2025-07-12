package mapper

import (
	"github.com/PCDattt/FintechEventProcessingSystem/server/internal/db"
	"github.com/PCDattt/FintechEventProcessingSystem/shared/enum"
	"github.com/PCDattt/FintechEventProcessingSystem/shared/model"
	"github.com/PCDattt/FintechEventProcessingSystem/shared/proto"
)

func DepositProtoToModel(tx *proto.DepositRequest) model.Transaction {
	toID := int(tx.ToAccountId)
	return model.Transaction{
		Type: enum.TransactionTypeDeposit,
		Status: enum.TransactionStatusPending,
		Amount: int(tx.Amount),
		Message: "Processing",
		ToAccountId: &toID,
		FromAccountId: nil,
	}
}

func WithdrawProtoToModel(tx *proto.WithdrawRequest) model.Transaction {
	fromID := int(tx.FromAccountId)
	return model.Transaction {
		Type: enum.TransactionTypeWithdraw,
		Status: enum.TransactionStatusPending,
		Amount: int(tx.Amount),
		Message: "Processing",
		ToAccountId: nil,
		FromAccountId: &fromID,
	}
}

func PaymentProtoToModel(tx *proto.PaymentRequest) model.Transaction {
	fromID := int(tx.FromAccountId)
	toID := int(tx.ToAccountId)

	return model.Transaction {
		Type: enum.TransactionTypePayment,
		Status: enum.TransactionStatusPending,
		Amount: int(tx.Amount),
		Message: "Processing",
		ToAccountId: &fromID,
		FromAccountId: &toID,
	}
}

func TransactionModelToCreateParams(tx model.Transaction) db.CreateTransactionParams {
	return db.CreateTransactionParams{
		Type: int32(tx.Type),
		Status: int32(enum.TransactionStatusPending),
		Amount: int32(tx.Amount),
		Message: "Processing",
		ToAccountID: tx.ToAccountId,
		FromAccountID: tx.FromAccountId,
	}
}

func DBTransactionToModel(tx db.CreateTransactionRow) model.Transaction {
	return model.Transaction{
		Id: int(tx.ID),
		Status: enum.TransactionStatus(tx.Status),
		Message: tx.Message,
	}
}