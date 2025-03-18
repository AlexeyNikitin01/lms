package grpc

import (
	"context"

	"github.com/pkg/errors"
)

func (s gRPCServerStruct) AuthByLoginPassword(ctx context.Context, req *UserAuthRequest) (*UserAuthResponse, error) {
	_, token, err := s.domainUser.AuthByLoginPassword(ctx, req.Login, req.Password)
	if err != nil {
		return nil, errors.Wrap(err, "auth by login password")
	}

	return &UserAuthResponse{
		Access:  token.Token,
		Refresh: token.Refresh,
	}, nil
}
