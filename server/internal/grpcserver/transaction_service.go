package grpcserver

import (
	"context"
	"log"

	"github.com/PCDattt/FintechEventProcessingSystem/shared/proto"
)

type TransactionServiceServer struct {
	proto.UnimplementedTransactionServiceServer
}

func NewTransactionServiceServer() *TransactionServiceServer {
	return &TransactionServiceServer{}
}

func (s *TransactionServiceServer) Deposit(ctx context.Context, req *proto.DepositRequest) (*proto.DepositResponse, error) {
	log.Printf("Received Deposit: %+v\n", req)
	return &proto.DepositResponse{Status: proto.TransactionStatus_TRANSACTION_STATUS_PENDING},nil
}

func (s *TransactionServiceServer) Withdraw(ctx context.Context, req *proto.WithdrawRequest) (*proto.WithdrawResponse, error) {
	log.Printf("Received Withdraw: %+v\n", req)
	return &proto.WithdrawResponse{Status: proto.TransactionStatus_TRANSACTION_STATUS_PENDING}, nil
}

func (s *TransactionServiceServer) Payment(ctx context.Context, req *proto.PaymentRequest) (*proto.PaymentResponse, error) {
	log.Printf("Received Payment: %+v\n", req)
	return &proto.PaymentResponse{Status: proto.TransactionStatus_TRANSACTION_STATUS_PENDING}, nil
}