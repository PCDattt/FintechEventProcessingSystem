package main

import (
	"context"
	"log"

	"github.com/PCDattt/FintechEventProcessingSystem/server/internal/db"
	"github.com/PCDattt/FintechEventProcessingSystem/server/internal/rabbitmq"
	"github.com/PCDattt/FintechEventProcessingSystem/server/internal/service"
	"github.com/PCDattt/FintechEventProcessingSystem/shared/config"
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
	transactionService := service.NewTransactionService(pool, queries)

	rabbitMQConsumer, err := rabbitmq.NewConsumer(cfg.RabbitURL, cfg.TransactionQueueName)
	if err != nil {
		log.Fatal("cannot create consumer:", err)
	}
	ctx := context.Background()
	
	err = rabbitMQConsumer.StartConsuming(ctx, transactionService.ProcessTransaction)
	if err != nil {
		log.Fatal("cannot create consumer:", err)
	}
	log.Println("Consuming transaction")
	select {}
}