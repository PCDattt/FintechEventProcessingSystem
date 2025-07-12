package grpcserver

import (
	"context"
	"log"

	"github.com/PCDattt/FintechEventProcessingSystem/server/internal/db"
	"github.com/PCDattt/FintechEventProcessingSystem/server/internal/mapper"
	"github.com/PCDattt/FintechEventProcessingSystem/server/internal/publisher"
	"github.com/PCDattt/FintechEventProcessingSystem/shared/enum"
	"github.com/PCDattt/FintechEventProcessingSystem/shared/proto"
)

type TransactionServiceServer struct {
	proto.UnimplementedTransactionServiceServer
	q *db.Queries
	p *publisher.Publisher
}

func NewTransactionServiceServer(q *db.Queries, p *publisher.Publisher) *TransactionServiceServer {
	return &TransactionServiceServer{
		q: q,
		p: p,
	}
}

func (s *TransactionServiceServer) SendTransaction(ctx context.Context, req *proto.TransactionRequest) (*proto.TransactionResponse, error) {
	model := mapper.TransactionProtoToModel(req)
	model.Status = enum.TransactionStatusProcessing
	model.Message = "Processing"
	params := mapper.TransactionModelToCreateParams(model)

	dbTransaction, err := s.q.CreateTransaction(ctx, params)
	if err != nil {
		log.Printf("Cannot create transaction in database: %v", err)
		return &proto.TransactionResponse {
			Status: proto.TransactionStatus_TRANSACTION_STATUS_FAILED,
			Message: "Cannot create transaction",
		}, err
	}

	model = mapper.DBTransactionToModel(dbTransaction)

	if err := s.p.PublishTransaction(model); err != nil {
		log.Printf("Cannot send transaction to RabbitMQ: %v", err)
		return &proto.TransactionResponse {
			Status: proto.TransactionStatus_TRANSACTION_STATUS_FAILED,
			Message: "Cannot process transaction",
		}, err
	}

	return &proto.TransactionResponse{
		Status: proto.TransactionStatus(dbTransaction.Status),
		Message: "Processing transaction",
		}, nil
}