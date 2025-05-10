package httpgin

import (
	"net/http"

	"lab/internal/app"

	"github.com/gin-gonic/gin"
)

func ping(_ *app.Lab) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Status(http.StatusOK)
		ctx.JSON(http.StatusOK, &gin.H{})
	}
}
