package app

import (
	"context"
	"testing"

	"edu-material/userMicroservice/internal/user"

	"github.com/stretchr/testify/mock"
)

func TestCreateUser(t *testing.T) {
	ctx := context.Background()
	r := &RepoUser{}
	r.On("CreateUser", ctx, user.User{}).Return(1, nil)
	a := CreateAppUser(r)

	u, err := a.CreateUser(ctx, user.User{}) // заполнить все поля!

	if u.ID != 1 {
		t.Error("error id is not correct")
	}

	if err != nil {
		t.Error(err)
	}
}

func TestDeleteUser(t *testing.T) {
	r := &RepoUser{}
	r.On("DeleteUser", mock.Anything, 1).Return(user.User{ID: 1}, nil)
	a := CreateAppUser(r)

	u, err := a.DeleteUser(context.Background(), 1)

	if u.ID != 1 {
		t.Error("error delete - not correct id")
	}

	if err != nil {
		t.Error(err)
	}
}

func TestUpdateUser(t *testing.T) {
	r := &RepoUser{}
	r.On("UpdateUser", mock.Anything, user.User{ID: 2}).Return(user.User{ID: 2}, nil)
	a := CreateAppUser(r)

	u, err := a.UpdateUser(context.Background(), user.User{ID: 2})

	if u.ID != 2 {
		t.Error("error update - not correct id")
	}

	if err != nil {
		t.Error(err)
	}
}

func TestGetUser(t *testing.T) {
	r := &RepoUser{}
	r.On("GetUser", mock.Anything, 1).Return(user.User{ID: 1}, nil)
	a := CreateAppUser(r)

	u, err := a.GetUser(context.Background(), 1)

	if u.ID != 1 {
		t.Error("error get - not correct id")
	}

	if err != nil {
		t.Error(err)
	}
}
