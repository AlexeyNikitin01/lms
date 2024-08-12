package httpgin

import (
	"log"

	"lms-user/internal/app"

	"github.com/gin-gonic/gin"
)

func ping(_ app.IAppUser) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Println("ping http gin")
		ctx.Status(200)
		ctx.JSON(200, &gin.H{})
	}
}
