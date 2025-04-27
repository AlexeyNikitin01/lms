package grpc

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"lms-user/internal/repository/pg/entity"
)

func (s gRPCServerStruct) GetUser(ctx context.Context, req *UserRequest) (*UserResponse, error) {
	user, err := s.domainUser.GetUser(ctx, req.Uuid)
	if err != nil {
		return nil, err
	}

	return &UserResponse{
		Uuid:        req.Uuid,
		Login:       user.Login,
		Name:        user.Name,
		Surname:     user.Surname,
		Email:       user.Email,
		Phone:       user.Phone,
		PlaceWork:   user.PlaceWork,
		Position:    user.Position,
		CreatedDate: timestamppb.New(user.CreatedAt),
	}, nil
}

func (s gRPCServerStruct) GetAllUser(ctx context.Context, _ *emptypb.Empty) (*UsersResponse, error) {
	users, err := s.domainUser.GetUsers(ctx)
	if err != nil {
		return nil, err
	}

	resp := make([]*UserResponse, 0, len(users))
	for _, user := range users {
		resp = append(resp, &UserResponse{
			Uuid:        user.ID,
			Login:       user.Login,
			Email:       user.Email,
			CreatedDate: timestamppb.New(user.CreatedAt),
		})
	}

	log.Println("get all user: ", len(users))

	return &UsersResponse{Users: resp}, nil
}

func Transform[T1 any, T2 any](s1 []T1, f func(a T1) T2) []T2 {
	if len(s1) == 0 {
		return []T2{}
	}

	s2 := make([]T2, 0, len(s1))

	for _, i := range s1 {
		s2 = append(s2, f(i))
	}

	return s2
}

func (s gRPCServerStruct) GetUserInfo(ctx context.Context, _ *emptypb.Empty) (*UserResponse, error) {
	fmt.Println("here")
	user := ctx.Value(UserCtxKey).(*entity.User)

	return &UserResponse{
		Uuid:        user.ID,
		Login:       user.Login,
		Name:        user.Name,
		Surname:     user.Surname,
		Email:       user.Email,
		Phone:       user.Phone,
		PlaceWork:   user.PlaceWork,
		Position:    user.Position,
		CreatedDate: timestamppb.New(user.CreatedAt),
	}, nil
}
