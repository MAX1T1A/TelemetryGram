package users

type UserAuthenticateResponse struct {
	ID   int    `json:"id"`
	Role string `json:"role"`
}
