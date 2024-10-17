package httpgin

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"course/internal/app"
)

func ping(_ app.ICourseApp) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Status(http.StatusOK)
		ctx.JSON(http.StatusOK, &gin.H{})
	}
}
