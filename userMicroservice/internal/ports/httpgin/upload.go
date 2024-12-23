package httpgin

import (
	"encoding/base64"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"

	"lms-user/internal/app"
)

func uploadAvatar(a app.IAppUser) gin.HandlerFunc {
	return func(c *gin.Context) {
		file, header, err := c.Request.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Не удалось получить файл: " + err.Error()})
			return
		}

		defer file.Close()

		allowedExtensions := map[string]bool{
			".jpg": true,
			".png": true,
		}

		ext := filepath.Ext(header.Filename)
		if !allowedExtensions[ext] {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Недопустимый тип файла"})
			return
		}

		currentUser := FromContext(c)

		active, url, err := a.UploadPhoto(c, file, header, currentUser.ID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Файл не загружен"})
			return
		}

		if active {
			c.JSON(http.StatusOK, gin.H{
				"message": "Файл успешно загружен",
				"url":     url,
			})
			return
		}

		buf, err := os.ReadFile(url)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Файл не загружен"})
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Файл успешно загружен локально",
			"img":     base64.StdEncoding.EncodeToString(buf),
		})
	}
}
