package model

import (
	"time"
	"github.com/PCDattt/FintechEventProcessingSystem/shared/enum"
)

type Transaction struct {
	Id int `db:"id" json:"id"`
	Type enum.TransactionType `db:"type" json:"type"`
	Status enum.TransactionStatus `db:"status" json:"status"`
	Amount int `db:"amount" json:"amount"`
	Message string `db:"message" json:"message"`
	AccountId int `db:"account_id" json:"accountId"`
	ToAccountId *int `db:"to_account_id" json:"toAccountId"`
	CreatedDate time.Time `db:"created_date" json:"createdDate"`
}