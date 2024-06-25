package httpgin

import (
	"edu-material/userMicroservice/internal/app"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func getAllUser(a app.IAppUser) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Printf("user")
		time.Sleep(time.Microsecond * 100)
		ctx.Status(200)
		ctx.JSON(200, &gin.H{"data": "ok"})
	}
}
