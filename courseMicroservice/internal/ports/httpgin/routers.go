package httpgin

import (
	"github.com/gin-gonic/gin"

	"course/internal/app"
)

func AppRouter(r *gin.RouterGroup, app app.ICourseApp) {
	r.GET("ping", ping(app))

	r.POST("add-course", addCourse(app))
	r.POST("add-lecture", addLecture(app))
}
