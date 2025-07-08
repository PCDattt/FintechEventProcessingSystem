package service

import (
	"context"

	"github.com/PCDattt/FintechEventProcessingSystem/server/internal/db"
	"github.com/PCDattt/FintechEventProcessingSystem/server/internal/mapper"
	"github.com/PCDattt/FintechEventProcessingSystem/shared/model"
)

type AccountService interface {
	CreateAccount(ctx context.Context, account model.Account) (model.Account, error)
}

type accountService struct {
	q *db.Queries
}

func NewAccountService(q *db.Queries) AccountService {
	return &accountService{q: q}
}

func (s *accountService) CreateAccount(ctx context.Context, account model.Account) (model.Account, error) {
	createAccountParams := mapper.ModelAccountToCreateParams(account)
	dbAccount, err := s.q.CreateAccount(ctx, createAccountParams)
	if err != nil {
		return model.Account{}, err
	}
	return mapper.DBAccountToModel(dbAccount), nil
}


