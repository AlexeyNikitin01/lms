package httpgin

import (
	"net/http"

	"lms-user/internal/app"

	"github.com/gin-gonic/gin"
)

func ping(_ app.IAppUser) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Status(http.StatusOK)
		ctx.JSON(http.StatusOK, &gin.H{})
	}
}
