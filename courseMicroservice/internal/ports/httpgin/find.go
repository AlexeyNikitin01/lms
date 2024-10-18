package httpgin

import (
	"net/http"

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
