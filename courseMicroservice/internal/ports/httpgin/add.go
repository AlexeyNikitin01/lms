package httpgin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"course/internal/app"
)

func addCourse(app app.ICourseApp) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req CourseRequest
		fmt.Println("start")

		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		fileContent, err := file.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open file"})
			return
		}
		defer fileContent.Close()

		data := make([]byte, file.Size)
		if _, err := fileContent.Read(data); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file"})
			return
		}

		url, err := app.UploadPhoto(c, data, file.Filename, "image/jpeg")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if err = c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		course, err := app.AddCourse(c, req.Name, req.Description, url)
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
