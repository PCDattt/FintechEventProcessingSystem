package router

import (
	"net/http"
	"github.com/PCDattt/FintechEventProcessingSystem/server/internal/handler"
)

func NewRouter(accountHandler *handler.AccountHandler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/account", accountHandler.CreateAccount)
	return mux
}