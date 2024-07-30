package grpc

import (
	"context"

	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/timestamppb"

	"lms-user/internal/repository/pg/entity"
)

func (s gRPCServerStruct) UpdateUser(ctx context.Context, req *UserUpdateRequest) (*UserResponse, error) {
	userForUpdate := entity.User{
		ID:        req.Uuid,
		Login:     req.Login,
		Name:      req.Name,
		Surname:   req.Surname,
		Phone:     req.Phone,
		Position:  req.Position,
		PlaceWork: req.PlaceWork,
		Email:     req.Email,
	}

	updateUser, err := s.domainUser.UpdateUser(ctx, &userForUpdate)
	if err != nil {
		return nil, errors.Wrap(err, "grpc interceptor: failed to update user")
	}

	return &UserResponse{
		Uuid:        updateUser.ID,
		Login:       updateUser.Login,
		Name:        updateUser.Name,
		Surname:     updateUser.Surname,
		Email:       updateUser.Email,
		Phone:       updateUser.Phone,
		PlaceWork:   updateUser.PlaceWork,
		Position:    updateUser.Position,
		CreatedDate: timestamppb.New(updateUser.CreatedAt),
	}, nil
}
