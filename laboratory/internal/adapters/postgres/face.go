package postgres

import (
	"fmt"
	"sync"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var once sync.Once

// LabPostgres паттерн одиночка.
type LabPostgres struct {
	DB *sqlx.DB
}

func NewLabPostgres(db *sqlx.DB) *LabPostgres {
	return &LabPostgres{DB: db}
}

func CreatePostgres(c *Config) (*sqlx.DB, error) {
	return sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		c.Host, c.Port, c.User, c.DBName, c.Password, c.SSLmode))
}
