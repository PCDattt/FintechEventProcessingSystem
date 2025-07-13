package grpcserver

import (
	"context"
	"log"

	"github.com/PCDattt/FintechEventProcessingSystem/server/internal/mapper"
	"github.com/PCDattt/FintechEventProcessingSystem/server/internal/rabbitmq"
	"github.com/PCDattt/FintechEventProcessingSystem/server/internal/service"
	"github.com/PCDattt/FintechEventProcessingSystem/shared/proto"
)

type TransactionServiceServer struct {
	proto.UnimplementedTransactionServiceServer
	svc service.TransactionService
	p *rabbitmq.Publisher
}

func NewTransactionServiceServer(svc service.TransactionService, p *rabbitmq.Publisher) *TransactionServiceServer {
	return &TransactionServiceServer{
		svc: svc,
		p: p,
	}
}

func (s *TransactionServiceServer) SendTransaction(ctx context.Context, req *proto.TransactionRequest) (*proto.TransactionResponse, error) {
	model := mapper.TransactionProtoToModel(req)
	
	s.svc.CreateTransaction(ctx, model)

	if err := s.p.PublishTransaction(model); err != nil {
		log.Printf("Cannot send transaction to RabbitMQ: %v", err)
		return &proto.TransactionResponse {
			Status: proto.TransactionStatus_TRANSACTION_STATUS_FAILED,
			Message: "Cannot process transaction",
		}, err
	}

	return &proto.TransactionResponse{
		Status: proto.TransactionStatus(model.Status),
		Message: "Processing transaction",
		}, nil
}