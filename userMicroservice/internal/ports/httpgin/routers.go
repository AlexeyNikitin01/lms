package httpgin

import (
	"lms-user/internal/app"

	"github.com/gin-gonic/gin"
)

func AppRouter(r *gin.RouterGroup, a app.IAppUser) {
	r.GET("ping", ping(a))

	r.POST("auth", authByLoginPassword(a))
	r.POST("register", register(a))
}
