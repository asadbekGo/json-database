package models

type CreateUser struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type User struct {
	Id        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type UserGetListRequest struct {
	Offset int
	Limit  int
}

type UserGetListResponse struct {
	Count int
	Users []*User
}
