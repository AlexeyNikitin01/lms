package httpgin

import (
	"github.com/AlexeyNikitin01/lms-user/internal/app"

	"github.com/gin-gonic/gin"
)

func AppRouter(r *gin.RouterGroup, a app.IAppUser) {
	r.GET("ping", ping(a))

	r.POST("get-user", getUser(a))
	r.POST("get-users", getAllUser(a))
	r.POST("auth", authByLoginPassword(a))
	r.POST("register", register(a))
	r.POST("update-user", updateUser(a))
	r.POST("upload-avatar", uploadAvatar(a))
}
