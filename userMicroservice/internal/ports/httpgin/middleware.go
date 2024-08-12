package httpgin

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"lms-user/internal/app"
	"lms-user/internal/repository/pg/entity"
)

var methods = map[string]struct{}{
	"/user/authByLoginPassword": {},
	"/user/register":            {},
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
