package postgres

import (
	"fmt"

	"edu-material/userMicroservice/internal/user"

	"github.com/jmoiron/sqlx"
)

type Postgres struct {
	DB *sqlx.DB
}

func GetAllUser() ([]user.User, error) {
	_ = fmt.Sprintf("SELECT * FROM %s", userTable)
	return nil, nil
}
