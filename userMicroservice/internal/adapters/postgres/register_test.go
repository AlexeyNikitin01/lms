package postgres_test

import (
	"context"
	"log"
	"testing"

	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/AlexeyNikitin01/lms-user/internal/adapters/postgres"
)

func TestRegister(t *testing.T) {
	type args struct {
		login    string
		password string
		email    string
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "success test install",
			args: args{
				login:    "testLogin",
				password: "testPassword",
				email:    "testEmail",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()

			cfg := postgres.Config{
				Host:     "localhost",
				Port:     "7777",
				User:     "postgres",
				DBName:   "user",
				Password: "pass",
				SSLmode:  "disable",
			}

			db, err := postgres.CreatePostgres(&cfg)
			if err != nil {
				log.Fatal(err)
			}

			boil.SetDB(db)

			repo := postgres.CreateRepoUser(db)

			createUser, err := repo.RegisterDB(ctx, tt.args.login, tt.args.password, tt.args.email)
			if err != nil {
				t.Fatal(err)
			}

			getUser, err := repo.GetUserDB(ctx, createUser.ID)
			if err != nil {
				t.Fatal(err)
			}

			if createUser.ID != getUser.ID {
				t.Fatalf("want %v, get %v", createUser.ID, getUser.ID)
			}

			err = repo.DeleteById(ctx, createUser.ID)
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}
