package grpcserver

import (
	"context"

	"github.com/PCDattt/FintechEventProcessingSystem/server/internal/db"
	"github.com/PCDattt/FintechEventProcessingSystem/server/internal/mapper"
	"github.com/PCDattt/FintechEventProcessingSystem/shared/proto"
)

type TransactionServiceServer struct {
	proto.UnimplementedTransactionServiceServer
	q *db.Queries
}

func NewTransactionServiceServer(q *db.Queries) *TransactionServiceServer {
	return &TransactionServiceServer{q: q}
}

func (s *TransactionServiceServer) Deposit(ctx context.Context, req *proto.DepositRequest) (*proto.DepositResponse, error) {
	model := mapper.DepositProtoToModel(req)
	params := mapper.TransactionModelToCreateParams(model)
	dbTransaction, err := s.q.CreateTransaction(ctx, params)
	if err != nil {
		return &proto.DepositResponse{}, err
	}
	return &proto.DepositResponse{Status: proto.TransactionStatus(dbTransaction.Status)}, nil
}

func (s *TransactionServiceServer) Withdraw(ctx context.Context, req *proto.WithdrawRequest) (*proto.WithdrawResponse, error) {
	model := mapper.WithdrawProtoToModel(req)
	params := mapper.TransactionModelToCreateParams(model)
	dbTransaction, err := s.q.CreateTransaction(ctx, params)
	if err != nil {
		return &proto.WithdrawResponse{}, err
	}
	return &proto.WithdrawResponse{Status: proto.TransactionStatus(dbTransaction.Status)}, nil
}

func (s *TransactionServiceServer) Payment(ctx context.Context, req *proto.PaymentRequest) (*proto.PaymentResponse, error) {
	model := mapper.PaymentProtoToModel(req)
	params := mapper.TransactionModelToCreateParams(model)
	dbTransaction, err := s.q.CreateTransaction(ctx, params)
	if err != nil {
		return &proto.PaymentResponse{}, err
	}
	return &proto.PaymentResponse{Status: proto.TransactionStatus(dbTransaction.Status)}, nil
}