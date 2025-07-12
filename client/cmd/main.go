package main

import (
	"context"
	"fmt"
	"log"

	"github.com/PCDattt/FintechEventProcessingSystem/client/internal/service"
	"github.com/PCDattt/FintechEventProcessingSystem/shared/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := proto.NewTransactionServiceClient(conn)
	txService := service.NewTransactionService(client)

	ctx := context.Background()

	depositReq := proto.TransactionRequest {
		AccountId: 1,
		Type: proto.TransactionType_TRANSACTION_TYPE_DEPOSIT,
		Amount: 1000,
	}

	withdrawReq := proto.TransactionRequest {
		AccountId: 1,
		Type: proto.TransactionType_TRANSACTION_TYPE_WITHDRAW,
		Amount: 500,
	}

	paymentReq := proto.TransactionRequest {
		AccountId: 1,
		Type: proto.TransactionType_TRANSACTION_TYPE_PAYMENT,
		ToAccountId: wrapperspb.Int32(5),
		Amount: 500,
	}

	depositRes, err := txService.SendTransaction(ctx, &depositReq)
	if err != nil {
		log.Fatalf("Deposit failed: %v\n", err)
	} else {
		fmt.Printf("Deposit status: %v\n", depositRes.Status)
	}

	withdrawRes, err := txService.SendTransaction(ctx, &withdrawReq)
	if err != nil {
		log.Fatalf("Withdraw failed: %v\n", err)
	} else {
		fmt.Printf("Withdraw status: %v\n", withdrawRes.Status)
	}

	paymentRes, err := txService.SendTransaction(ctx, &paymentReq)
	if err != nil {
		log.Fatalf("Payment failed: %v\n", err)
	} else {
		fmt.Printf("Payment status: %v\n", paymentRes.Status)
	}
}