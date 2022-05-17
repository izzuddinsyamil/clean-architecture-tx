package handler

type GetUserRequest struct {
	Id string `param:"id"`
}

type CreateUserRequest struct {
	Name    string `json:"name"`
	Balance int    `json:"balance"`
}
