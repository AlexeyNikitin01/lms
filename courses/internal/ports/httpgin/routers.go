package httpgin

import (
	"github.com/gin-gonic/gin"

	"course/internal/app"
)

func AppRouter(r *gin.RouterGroup, app app.ICourseApp) {
	r.GET("ping", ping(app))
	r.GET(":id/users", getListUserByCourseID(app))
	r.GET(":id/user/:uuid", getRole(app))

	//TODO: заменить на get запрос, через query параметр.
	r.POST("all", getAllCourses(app))
	r.POST("get/:id", getCourse(app))
	r.POST("create", createCourse(app))
	r.POST("update/:id", updateCourse(app))
}
