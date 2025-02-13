package postgres

import (
	"context"

	"github.com/jmoiron/sqlx"

	"lms-user/internal/repository/pg/entity"
)

type IUserPostgres interface {
	RegisterDB(ctx context.Context, login, password, email string) (*entity.User, error)
	GetUserDB(ctx context.Context, uuid string) (*entity.User, error)
	GetUserByLoginPasswordDB(ctx context.Context, login, password string) (*entity.User, error)
	GetUserByTokenDB(ctx context.Context, tokenID string) (*entity.User, *entity.Token, error)
	InsertTokenDB(ctx context.Context, token *entity.Token) error
	UpdateUser(ctx context.Context, user *entity.User) (*entity.User, error)
	GetTokenByUserID(ctx context.Context, userID string) (*entity.Token, error)
	DeleteOldToken(ctx context.Context, userID string) error
	DeleteById(ctx context.Context, userID string) error
	SaveAvatarFileName(ctx context.Context, url, userID string) error
	GetUsersDB(ctx context.Context) (entity.UserSlice, error)
}

/*

	Для использования sql-boiler необходимо использовать драйвер DB
	или boil.GetContextDB() на выбор

*/

type RepoUser struct {
	DB *sqlx.DB
}

func CreateRepoUser(db *sqlx.DB) IUserPostgres {
	return &RepoUser{
		DB: db,
	}
}
