package service

import (
	"context"

	"github.com/PCDattt/FintechEventProcessingSystem/shared/proto"
)

type TransactionServiceInterface interface {
	Deposit(context.Context, *proto.DepositRequest) (*proto.DepositResponse, error)
	Withdraw(context.Context, *proto.WithdrawRequest) (*proto.WithdrawResponse, error)
	Payment(context.Context, *proto.PaymentRequest) (*proto.PaymentResponse, error)
}

type TransactionService struct {
	client proto.TransactionServiceClient
}

func NewTransactionService(client proto.TransactionServiceClient) *TransactionService {
	return &TransactionService{client: client}
}

func (s *TransactionService) Deposit(ctx context.Context, req *proto.DepositRequest) (*proto.DepositResponse, error) {
	res, err := s.client.Deposit(ctx, req)
	if err != nil {
		return &proto.DepositResponse{}, err
	}
	return res, nil
}

func (s *TransactionService) Withdraw(ctx context.Context, req *proto.WithdrawRequest) (*proto.WithdrawResponse, error) {
	res, err := s.client.Withdraw(ctx, req)
	if err != nil {
		return &proto.WithdrawResponse{}, err
	}
	return res, nil
}


func (s *TransactionService) Payment(ctx context.Context, req *proto.PaymentRequest) (*proto.PaymentResponse, error) {
	res, err := s.client.Payment(ctx, req)
	if err != nil {
		return &proto.PaymentResponse{}, err
	}
	return res, nil
}