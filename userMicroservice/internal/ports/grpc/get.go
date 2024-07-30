package grpc

import (
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"
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
