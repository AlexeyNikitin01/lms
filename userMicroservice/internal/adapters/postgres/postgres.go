package postgres

import (
	"fmt"
	"sync"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var once sync.Once

type UserPostgres struct {
	db  *sqlx.DB
	err error
}

var instance *UserPostgres

func GetInstance() *UserPostgres {
	once.Do(
		func() {
			instance = new(UserPostgres)
		},
	)
	return instance
}

func CreatePostgres(c *Config) (*sqlx.DB, error) {
	userPostgres := GetInstance()
	userPostgres.db, userPostgres.err = sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		c.Host, c.Port, c.User, c.DBName, c.Password, c.SSLmode))
	if userPostgres.err != nil {
		return nil, userPostgres.err
	}

	userPostgres.err = userPostgres.db.Ping()
	if userPostgres.err != nil {
		return nil, userPostgres.err
	}

	return userPostgres.db, nil
}
