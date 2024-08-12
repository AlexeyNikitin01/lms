package httpgin

import (
	"lms-user/internal/app"

	"github.com/gin-gonic/gin"
)

func ping(_ app.IAppUser) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Status(200)
		ctx.JSON(200, &gin.H{})
	}
}
