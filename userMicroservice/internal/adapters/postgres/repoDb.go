package postgres

import (
	"context"

	"edu-material/userMicroservice/internal/user"

	"github.com/jmoiron/sqlx"
)

type RepoUser struct {
	DB *sqlx.DB
}

func(r RepoUser) CreateUser(ctx context.Context, user user.User) (int, error) {return -1, nil}
func(r RepoUser) DeleteUser(ctx context.Context, id int) (*user.User, error) {return nil, nil} 
func(r RepoUser) GetUser(ctx context.Context, id int) (*user.User, error) {return nil, nil}
func(r RepoUser) UpdateUser(ctx context.Context, user user.User) (*user.User, error) {return nil, nil}

func CreateRepoUser(db *sqlx.DB) *RepoUser {
	return &RepoUser{
		DB: db,
	}
}
