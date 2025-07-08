package response

type CreateAccountResponse struct {
	Id int `json:"id"`
	Username string `json:"username"`
	Amount int `json:"amount"`
}