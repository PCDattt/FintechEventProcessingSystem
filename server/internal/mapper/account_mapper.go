package mapper

import (
	"github.com/PCDattt/FintechEventProcessingSystem/shared/model"
	"github.com/PCDattt/FintechEventProcessingSystem/server/internal/db"
	"github.com/PCDattt/FintechEventProcessingSystem/server/internal/handler/request"
	"github.com/PCDattt/FintechEventProcessingSystem/server/internal/handler/response"
)

func DBAccountToModel(a db.Account) model.Account {
	return model.Account{
		Id: int(a.ID),
		Username: a.Username,
		Password: a.Password,
		Amount: int(a.Amount),
		CreatedDate: a.CreatedDate,
	}
}

func ModelAccountToCreateParams(a model.Account) db.CreateAccountParams {
	return db.CreateAccountParams{
		Username: a.Username,
		Password: a.Password,
	}
}

func CreateAccountRequestToModel(rq request.CreateAccountRequest) model.Account {
	return model.Account{
		Username: rq.Username,
		Password: rq.Password,
	}
}

func ModelAccountToCreateResponse(a model.Account) response.CreateAccountResponse {
	return response.CreateAccountResponse{
		Id: a.Id,
		Username: a.Username,
		Amount: a.Amount,
	}
}