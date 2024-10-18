package httpgin

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"course/internal/app"
)

func addCourse(app app.ICourseApp) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req CourseRequest
		err := c.BindJSON(&req)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		course, err := app.AddCourse(c, req.Name)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"course": course})
	}
}

func addLecture(app app.ICourseApp) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req LectureRequest
		err := c.BindJSON(&req)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = app.AddLecture(c, req.Title, req.Lecture, req.CourseID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"course_id": req.CourseID})
	}
}
