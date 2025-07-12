package service

import (
	"context"

	"github.com/PCDattt/FintechEventProcessingSystem/shared/proto"
)

type TransactionServiceInterface interface {
	SendTransaction(context.Context, *proto.TransactionRequest) (*proto.TransactionResponse, error)
}

type TransactionService struct {
	client proto.TransactionServiceClient
}

func NewTransactionService(client proto.TransactionServiceClient) *TransactionService {
	return &TransactionService{client: client}
}

func (s *TransactionService) SendTransaction(ctx context.Context, req *proto.TransactionRequest) (*proto.TransactionResponse, error) {
	res, err := s.client.SendTransaction(ctx, req)
	if err != nil {
		return &proto.TransactionResponse{}, err
	}
	return res, nil
}