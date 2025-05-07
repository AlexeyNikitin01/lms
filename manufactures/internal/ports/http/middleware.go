package httpgin

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"

	user "manufactures/pkg/grpc"
)

var methods = map[string]struct{}{
	"/metrics": {},
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

type contextKey struct {
	name string
}

var UserCtxKey = &contextKey{name: "user"}

func auth(client user.UserServiceClient) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if _, ok := methods[ctx.FullPath()]; ok {
			ctx.Next()
			return
		}

		if err := pingToUser(ctx, 10*time.Second, 3, client); err != nil {
			return
		}

		path := ctx.GetHeader("authorization")

		user, err := client.GetUserInfo(metadata.NewOutgoingContext(ctx, metadata.New(map[string]string{
			"authorization": path,
		})), &emptypb.Empty{})
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"auth": "error", "msg": err.Error()})
			return
		}

		ctx.Set(UserCtxKey.name, user)

		ctx.Next()
	}
}

func pingToUser(ctx context.Context, timeout time.Duration, times int, client user.UserServiceClient) error {
	times--
	if _, err := client.Ping(ctx, &emptypb.Empty{}); err == nil {
		return nil
	}

	// Иначе пробуем достучаться
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	ticker := time.NewTicker(timeout / time.Duration(times))
	defer ticker.Stop()

	attemp := 0

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			attemp++

			if _, err := client.Ping(ctx, &emptypb.Empty{}); err != nil {
				if attemp > times {
					return errors.New("timeout")
				}
			} else if err == nil {
				return nil
			}
		}
	}
}
