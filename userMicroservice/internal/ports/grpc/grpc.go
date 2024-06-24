package grpc

import (
	"context"
	"edu-material/userMicroservice/internal/app"
	"edu-material/userMicroservice/internal/repository/pg/entity"

	"github.com/friendsofgo/errors"
	"github.com/google/uuid"
	authGrpc "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	"github.com/pascaldekloe/jwt"
	"github.com/thanhpk/randstr"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"strings"
	"time"
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
		Uuid:  req.Uuid,
		Login: user.Login,
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
	user, err := entity.Users(
		entity.UserWhere.Login.EQ(req.Login),
		entity.UserWhere.Password.EQ([]byte(req.Password)),
	).One(ctx, boil.GetContextDB())
	if err != nil {
		return nil, errors.Wrap(err, "AuthByLoginPassword - not found user")
	}

	tokensUser := &entity.Token{
		ID:     uuid.New().String(),
		UserID: user.ID,
	}

	access, err := accessToken(user, tokensUser)
	if err != nil {
		return nil, err
	}

	refresh, err := refreshToken(tokensUser)
	if err != nil {
		return nil, err
	}
	//todo: происходит задубливание в бд, для одного пользователя может быть много токенов
	err = tokensUser.Insert(ctx, boil.GetContextDB(), boil.Infer())
	if err != nil {
		return nil, errors.Wrap(err, "AuthByLoginPassword - failed to insert token")
	}

	return &UserAuthResponse{
		Access:  access,
		Refresh: refresh,
	}, nil
}

func refreshToken(tokenUser *entity.Token) (string, error) {
	refresh := randstr.String(30)

	tokenUser.Refresh = refresh

	return refresh, nil
}

func accessToken(user *entity.User, tokenUser *entity.Token) (string, error) {
	now := time.Now()
	accessExpires := now.Add(time.Hour * 1000)

	claims := jwt.Claims{
		Registered: jwt.Registered{
			Issued:  jwt.NewNumericTime(now.Truncate(time.Second)),
			Expires: jwt.NewNumericTime(accessExpires.Truncate(time.Second)),
			Subject: tokenUser.ID,
		},
		Set: map[string]interface{}{
			"user_id": user.ID,
		},
	}

	token, err := claims.HMACSign(jwt.HS512, []byte("00000000"))
	if err != nil {
		return "", err
	}

	tokenUser.Token = string(token)
	tokenUser.ExpiresAt = accessExpires

	return string(token), nil
}

func parseToken(ctx context.Context, token string) (*entity.User, *entity.Token, error) {
	clm, err := jwt.HMACCheck([]byte(token), []byte("00000000"))
	if err != nil {
		return nil, nil, err
	}

	if !clm.Valid(time.Now()) {
		return nil, nil, err
	}

	tokenUserId := clm.Subject

	tokenUser, err := entity.Tokens(entity.TokenWhere.ID.EQ(tokenUserId)).One(ctx, boil.GetContextDB())
	if err != nil {
		return nil, nil, err
	}

	user, err := entity.Users(entity.UserWhere.ID.EQ(tokenUser.UserID)).One(ctx, boil.GetContextDB())
	if err != nil {
		return nil, nil, err
	}

	return user, tokenUser, nil
}

func Interceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		return authGrpc.UnaryServerInterceptor(authFunc())(ctx, req, info, handler)
	}
}

func authFunc() authGrpc.AuthFunc {
	return func(ctx context.Context) (context.Context, error) {
		token, err := FromMD(ctx, "bearer")
		if err != nil {
			return nil, errors.Wrap(err, "grpc interceptor: failed to parse bearer token")
		}

		user, _, err := parseToken(ctx, token)
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

func NewService(domainUser app.IAppUser) Server {
	return &gRPCServerStruct{
		domainUser: domainUser,
	}
}
