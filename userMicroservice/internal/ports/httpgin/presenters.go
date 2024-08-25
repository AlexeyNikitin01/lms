package httpgin

import (
	"time"
)

type UserResponse struct {
	UUID        string    `json:"uuid"`
	Login       string    `json:"login"`
	Name        string    `json:"name"`
	Surname     string    `json:"surname"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
	PlaceWork   string    `json:"place_work"`
	Position    string    `json:"position"`
	CreatedDate time.Time `json:"created_date"`
}

type UserRegisterResponse struct {
	UUID  string `json:"uuid"`
	Login string `json:"login"`
	Email string `json:"email"`
}

type UserRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Error struct{}

type AuthByLogPassRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type AuthByLogPassResponse struct {
	Token    string `json:"token"`
	Refresh  string `json:"refresh"`
	UserUUID string `json:"user_uuid"`
}

type GetUserReq struct {
	UUID string `json:"user_uuid"`
}

type UpdateUserRequest struct {
	UUID      string `json:"user_uuid"`
	Login     string `json:"login"`
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	PlaceWork string `json:"place_work"`
	Position  string `json:"position"`
}
