package grpc

import (
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
	"lms-user/internal/app"
	"lms-user/internal/repository/pg/entity"

	"github.com/friendsofgo/errors"
	authGrpc "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"strings"
)

type contextKey struct {
	name string
}

var UserCtxKey = &contextKey{name: "user"}

type Server interface {
	UserServiceServer
}

type gRPCServerStruct struct {
	UnimplementedUserServiceServer
	domainUser app.IAppUser
}

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

func (s gRPCServerStruct) Ping(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, nil
}

func (s gRPCServerStruct) RegisterUser(ctx context.Context, req *UserRegisterRequest) (*UserRegisterResponse, error) {
	u, err := s.domainUser.Register(ctx, req.Login, req.Password)
	if err != nil {
		return nil, err
	}

	return &UserRegisterResponse{
		Uuid:  u.ID,
		Login: u.Login,
	}, nil
}

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

func (s gRPCServerStruct) Interceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		return authGrpc.UnaryServerInterceptor(s.authFunc())(ctx, req, info, handler)
	}
}

func (s gRPCServerStruct) authFunc() authGrpc.AuthFunc {
	return func(ctx context.Context) (context.Context, error) {
		token, err := FromMD(ctx, "bearer")
		if err != nil {
			return nil, errors.Wrap(err, "grpc interceptor: failed to parse bearer token")
		}

		user, _, err := s.domainUser.ParseToken(ctx, token)
		if err != nil {
			return nil, errors.Wrap(err, "grpc interceptor: failed to parse token")
		}

		return context.WithValue(ctx, UserCtxKey, user), nil
	}
}

func FromMD(ctx context.Context, key string) (string, error) {
	val := metautils.ExtractIncoming(ctx).Get("authorization")
	if val == "" {
		return "", status.New(codes.Unauthenticated, "expected bearer token").Err()
	}

	splits := strings.SplitN(val, " ", 2)
	if len(splits) != 2 {
		return "", status.New(codes.Unauthenticated, "invalid token").Err()
	}

	if strings.EqualFold(splits[0], key) {
		return "", status.New(codes.PermissionDenied, "permission denied").Err()
	}

	return splits[1], nil
}

func FromContext(ctx context.Context) *entity.User {
	user, ok := ctx.Value(UserCtxKey).(*entity.User)
	if !ok {
		return nil
	}

	return user
}

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

func NewService(domainUser app.IAppUser) Server {
	return &gRPCServerStruct{
		domainUser: domainUser,
	}
}
