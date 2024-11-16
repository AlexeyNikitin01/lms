package httpgin

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/AlexeyNikitin01/lms-user/internal/app"
	"github.com/AlexeyNikitin01/lms-user/internal/repository/pg/entity"
)

var methods = map[string]struct{}{
	"/metrics":        {},
	"/user/auth":      {},
	"/user/register":  {},
	"/user/get-users": {},
}

type contextKey struct {
	name string
}

var UserCtxKey = &contextKey{name: "user"}

func auth(app app.IAppUser) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if _, ok := methods[ctx.FullPath()]; ok {
			ctx.Next()
			return
		}

		path := ctx.GetHeader("authorization")

		tokenRaw := strings.ReplaceAll(path, "bearer ", "")

		user, _, err := app.ParseToken(ctx, tokenRaw)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"auth": "error", "msg": err.Error()})
			return
		}

		ctx.Set(UserCtxKey.name, user)

		ctx.Next()
	}
}

func FromContext(ctx context.Context) *entity.User {
	user, ok := ctx.Value(UserCtxKey.name).(*entity.User)
	if !ok {
		return nil
	}

	return user
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
