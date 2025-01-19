package httpgin

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"course/internal/app"
)

func getLecture(app app.ICourseApp) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req FindLecturesRequest
		err := c.BindJSON(&req)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		lecture, err := app.FindLecture(c, req.CourseID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"lecture": lecture})
	}
}

func getAllCourses(app app.ICourseApp) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req AllCoursesRequest
		err := c.BindJSON(&req)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		courses, total, err := app.AllCourse(c, req.Limit, req.Offset)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		for _, course := range courses {
			course.PhotoURL = "https://storage.yandexcloud.net/lms-user/" + course.PhotoURL
		}

		c.JSON(http.StatusOK, gin.H{"courses": courses, "total": total})
	}
}

func getCourse(app app.ICourseApp) gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		courseID, err := strconv.ParseInt(idParam, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
			return
		}

		course, err := app.GetCourse(c, courseID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		modules := []gin.H{}
		for _, m := range course.R.GetModules() {
			lectures := []gin.H{}
			for _, l := range m.R.GetLectures() {
				tests := []gin.H{}
				for _, t := range l.R.GetTests() {
					questions := []gin.H{}
					for _, q := range t.R.GetQuestions() {
						questions = append(questions, gin.H{
							"question": q,
						})
					}
					tests = append(tests, gin.H{
						"test":      t,
						"questions": questions,
					})

				}
				lectures = append(lectures, gin.H{
					"lecture": l,
					"tests":   tests,
				})
			}
			modules = append(modules, gin.H{
				"module":   m,
				"lectures": lectures,
			})

		}

		c.JSON(
			http.StatusOK,
			gin.H{
				"course":  course,
				"modules": modules,
			},
		)
	}
}
