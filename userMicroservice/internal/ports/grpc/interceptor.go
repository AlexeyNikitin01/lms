package grpc

import (
	"context"
	"strings"

	authGrpc "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/lms-user/internal/repository/pg/entity"
)

type contextKey struct {
	name string
}

var UserCtxKey = &contextKey{name: "user"}

var noAuthMethod = map[string]struct{}{
	"/user.UserService/authByLoginPassword": {},
	"/user.UserService/registerUser":        {},
	"/user.UserService/getAllUser":          {},
}

func (s gRPCServerStruct) Interceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		if _, ok := noAuthMethod[info.FullMethod]; ok {
			return handler(ctx, req)
		}

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

	if !strings.EqualFold(splits[0], key) {
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
