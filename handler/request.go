package handler

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type GetUserRequest struct {
	Id string `param:"id"`
}

type CreateUserRequest struct {
	Name    string `json:"name"`
	Balance int    `json:"balance"`
}

type TransactRequest struct {
	SenderId   int `json:"sender_id"`
	ReceiverId int `json:"receiver_id"`
	Amount     int `json:"amount"`
}
