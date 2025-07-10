package main

import (
	"context"
	"fmt"
	"log"

	"github.com/PCDattt/FintechEventProcessingSystem/client/internal/service"
	"github.com/PCDattt/FintechEventProcessingSystem/shared/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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

	depositReq := proto.DepositRequest {
		ToAccountId: 1,
		Amount: 1000,
	}

	withdrawReq := proto.WithdrawRequest {
		FromAccountId: 1,
		Amount: 500,
	}

	paymentReq := proto.PaymentRequest {
		FromAccountId: 1,
		ToAccountId: 2,
		Amount: 500,
	}

	depositRes, err := txService.Deposit(ctx, &depositReq)
	if err != nil {
		log.Fatalf("Deposit failed: %v\n", err)
	} else {
		fmt.Printf("Deposit status: %v\n", depositRes.Status)
	}

	withdrawRes, err := txService.Withdraw(ctx, &withdrawReq)
	if err != nil {
		log.Fatalf("Withdraw failed: %v\n", err)
	} else {
		fmt.Printf("Withdraw status: %v\n", withdrawRes.Status)
	}

	paymentRes, err := txService.Payment(ctx, &paymentReq)
	if err != nil {
		log.Fatalf("Payment failed: %v\n", err)
	} else {
		fmt.Printf("Payment status: %v\n", paymentRes.Status)
	}
}