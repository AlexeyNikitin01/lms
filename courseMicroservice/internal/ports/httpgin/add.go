package httpgin

import (
	"log"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"course/internal/app"
	"course/internal/repository/pg/entity"
)

func addCourse(app app.ICourseApp) gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.PostForm("name")
		description := c.PostForm("description")

		course, err := app.AddCourse(c, name, description)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if err = addAvatarCourse(app, c, course); err != nil {
			log.Println("аватар для курса не загружен")
		}

		c.JSON(http.StatusOK, gin.H{"course": course})
	}
}

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

func addAvatarCourse(app app.ICourseApp, c *gin.Context, course *entity.Course) error {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		return err
	}

	defer file.Close()

	allowedExtensions := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
	}

	if !allowedExtensions[filepath.Ext(header.Filename)] {
		return errors.New("unsupported file type, only .jpg, .jpeg, and .png are allowed")
	}

	if err = app.UploadPhoto(c, file, header, course); err != nil {
		return err
	}

	return nil
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
