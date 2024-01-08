package user

import (
	"context"
)

//go:generate mockery --output ../app --name mockRepoUser
type RepoUser interface {
	CreateUser(ctx context.Context, user User) (int, error)
	DeleteUser(ctx context.Context, id int) (*User, error)
	GetUser(ctx context.Context, id int) (*User, error)
	UpdateUser(ctx context.Context, user User) (*User, error)
}
