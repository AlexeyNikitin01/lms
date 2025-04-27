package httpgin

import (
	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"course/internal/app"
	"course/internal/repository/pg/entity"
	grpcuser "course/pkg/grpc-user"
)

func createCourse(app app.ICourseApp) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, ok := c.Get("user")
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{})
			return
		}

		castUser := user.(*grpcuser.UserResponse)

		name := c.PostForm("name")
		description := c.PostForm("description")

		course, err := app.AddCourse(c, name, description, castUser.GetUuid())
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
