package httpgin

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"course/internal/app"
)

func updateCourse(app app.ICourseApp) gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		courseID, err := strconv.ParseInt(idParam, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
			return
		}

		var course Course

		if err := c.ShouldBindJSON(&course); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err = app.UpdateCourse(c, courseID, convertToEntityCourse(course, courseID)); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"course": courseID})
	}
}
