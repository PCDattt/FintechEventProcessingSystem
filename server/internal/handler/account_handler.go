package handler

import (
	"net/http"
	"github.com/PCDattt/FintechEventProcessingSystem/server/internal/handler/request"
	"github.com/PCDattt/FintechEventProcessingSystem/server/internal/service"
	"github.com/PCDattt/FintechEventProcessingSystem/server/internal/mapper"
	"github.com/gin-gonic/gin"
	"github.com/PCDattt/FintechEventProcessingSystem/server/internal/prometheus"
)

type AccountHandler struct {
	svc service.AccountService
}

func NewAccountHandler(svc service.AccountService) *AccountHandler {
	return &AccountHandler{svc: svc}
}

func (h *AccountHandler) CreateAccount(c *gin.Context) {
	var req request.CreateAccountRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		prometheus.AccountRequestsTotal.WithLabelValues("CreateAccount", "400").Inc()
		return
	}
	
	account := mapper.CreateAccountRequestToModel(req)
	account, err := h.svc.CreateAccount(c, account)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		prometheus.AccountRequestsTotal.WithLabelValues("CreateAccount", "500").Inc()
		return
	}

	res := mapper.ModelAccountToCreateResponse(account)

	c.JSON(http.StatusOK, res)
	prometheus.AccountRequestsTotal.WithLabelValues("CreateAccount", "200").Inc()
}