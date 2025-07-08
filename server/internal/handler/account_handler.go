package handler

import (
	"encoding/json"
	"net/http"
	"github.com/PCDattt/FintechEventProcessingSystem/server/internal/handler/request"
	"github.com/PCDattt/FintechEventProcessingSystem/server/internal/service"
	"github.com/PCDattt/FintechEventProcessingSystem/server/internal/mapper"
)

type AccountHandler struct {
	svc service.AccountService
}

func NewAccountHandler(svc service.AccountService) *AccountHandler {
	return &AccountHandler{svc: svc}
}

func (h *AccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var req request.CreateAccountRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	
	account := mapper.CreateAccountRequestToModel(req)
	account, err := h.svc.CreateAccount(r.Context(), account)
	if err != nil {
		http.Error(w, "could not create account", http.StatusInternalServerError)
		return
	}

	res := mapper.ModelAccountToCreateResponse(account)

	json.NewEncoder(w).Encode(res)
}