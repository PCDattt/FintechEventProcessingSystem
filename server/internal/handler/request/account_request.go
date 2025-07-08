package request

type CreateAccountRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}