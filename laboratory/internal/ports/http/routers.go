package httpgin

import (
	"lab/internal/app"

	"github.com/gin-gonic/gin"
)

func AppRouter(r *gin.RouterGroup, a *app.Lab) {
	r.GET("ping", ping(a))
	r.GET(":id", getLab(a))
	r.PUT(":id", updateLab(a))

	r.GET("models/:id", getAirplaneModel(a))
}
