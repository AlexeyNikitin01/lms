package httpgin

import (
	"time"
)

type UserResponse struct {
	Uuid        string    `json:"uuid"`
	Login       string    `json:"login"`
	Name        string    `json:"name"`
	Surname     string    `json:"surname"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
	PlaceWork   string    `json:"place_work"`
	Position    string    `json:"position"`
	CreatedDate time.Time `json:"created_date"`
}

type UserRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Error struct {
}
