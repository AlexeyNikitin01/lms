package httpgin

import (
	"edu-material/userMicroservice/internal/app"

	"github.com/gin-gonic/gin"
)

func AppRouter(r *gin.RouterGroup, a app.IAppUser) {
	r.GET("all-user", getAllUser(a))
}
