package httpgin

import (
	"github.com/gin-gonic/gin"

	"course/internal/app"
)

func AppRouter(r *gin.RouterGroup, app app.ICourseApp) {
	r.GET("ping", ping(app))

	//TODO: заменить на get запрос, через query параметр.
	r.POST("get-lecture", getLecture(app))
	r.POST("get-all-courses", getAllCourses(app))
	r.POST("get-course/:id", getCourse(app))

	r.POST("add-course", addCourse(app))
	r.POST("update-course/:id", updateCourse(app))
	r.POST("add-lecture", addLecture(app))
}
