package model

import "time"

type Account struct {
	Id int `db:"id" json:"id"`
	Username string `db:"username" json:"username"`
	Password string `db:"password" json:"password"`
	Amount int `db:"amount" json:"amount"`
	CreatedDate time.Time `db:"created_date" json:"createdDate"`
}