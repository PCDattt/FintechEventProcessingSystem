package main

import (
	"context"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/PCDattt/FintechEventProcessingSystem/shared/config"
	"github.com/PCDattt/FintechEventProcessingSystem/server/internal/db"
	"github.com/PCDattt/FintechEventProcessingSystem/server/internal/handler"
	"github.com/PCDattt/FintechEventProcessingSystem/server/internal/router"
	"github.com/PCDattt/FintechEventProcessingSystem/server/internal/service"
)

func main() {
	cfg := config.LoadConfig()
	
	pool, err := pgxpool.New(context.Background(), cfg.DBURL)
	if err != nil {
		log.Fatal("cannot parse DB config:", err)
	}

	err = pool.Ping(context.Background())
	if err != nil {
		log.Fatal("cannot connect to DB:", err)
	}
	defer pool.Close()

	queries := db.New(pool)
	accountService := service.NewAccountService(queries)
	accountHandler := handler.NewAccountHandler(accountService)

	r := router.NewRouter(accountHandler)

	log.Println("Listening on: 8080")
	http.ListenAndServe(":8080", r)
}