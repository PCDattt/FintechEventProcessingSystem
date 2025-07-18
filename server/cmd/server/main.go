package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"

	"github.com/PCDattt/FintechEventProcessingSystem/server/internal/db"
	"github.com/PCDattt/FintechEventProcessingSystem/server/internal/grpcserver"
	"github.com/PCDattt/FintechEventProcessingSystem/server/internal/handler"
	"github.com/PCDattt/FintechEventProcessingSystem/server/internal/rabbitmq"
	"github.com/PCDattt/FintechEventProcessingSystem/server/internal/router"
	"github.com/PCDattt/FintechEventProcessingSystem/server/internal/service"
	"github.com/PCDattt/FintechEventProcessingSystem/shared/config"
	"github.com/PCDattt/FintechEventProcessingSystem/shared/proto"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	cfg := config.LoadConfig()
	
	pool, err := pgxpool.New(context.Background(), cfg.DBURL)
	if err != nil {
		log.Fatal("cannot parse DB config:", err)
	}
	defer pool.Close()

	err = pool.Ping(context.Background())
	if err != nil {
		log.Fatal("cannot connect to DB:", err)
	}

	queries := db.New(pool)
	accountService := service.NewAccountService(queries)
	accountHandler := handler.NewAccountHandler(accountService)
	transactionService := service.NewTransactionService(pool, queries)

	r := router.NewRouter(accountHandler)

	httpServer := &http.Server{
	Addr:    ":8080",
	Handler: r,
}

	log.Println("http server listening on: 8080")
	go func() {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTP server error: %v", err)
		}
	}()

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	publisher, err := rabbitmq.NewPublisher(cfg.RabbitURL, cfg.TransactionQueueName)
	if err != nil {
		log.Fatal("failed to connect to RabbitMQ:", err)
	}
	defer publisher.Close()


	proto.RegisterTransactionServiceServer(grpcServer, grpcserver.NewTransactionServiceServer(transactionService, publisher))

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	log.Println("gRPC server listening on: 50051")
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve grpc: %v", err)
		}
	}()

	<-ctx.Done()
	log.Println("Shutting down...")

	// Graceful shutdown of HTTP server
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(shutdownCtx); err != nil {
		log.Printf("HTTP server shutdown error: %v", err)
	}

	// Shutdown gRPC server
	grpcServer.GracefulStop()
	lis.Close()

	log.Println("Servers gracefully stopped")
}