package httpgin

import (
	"encoding/base64"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"lms-user/internal/app"
)

func getUser(a app.IAppUser) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req GetUserReq
		if err := c.Bind(&req); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"method": c.FullPath(),
				"err":    errors.Wrap(err, "get user request").Error(),
			})

			return
		}

		currentUser := FromContext(c)
		if !checkUser(currentUser, req.UUID) {
			c.JSON(http.StatusInternalServerError, gin.H{
				"method": c.FullPath(),
				"err":    errors.New("you don`t get user").Error(),
			})

			return
		}

		user, err := a.GetUser(c, req.UUID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"method": c.FullPath(),
				"err":    errors.Wrap(err, "get user domain").Error(),
			})

			return
		}

		buf := make([]byte, 0)
		if strings.Contains(user.Avatar, "tmp") {
			buf, err = os.ReadFile(user.Avatar)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Файл не загружен"})
				return
			}
			user.Avatar = ""
		}

		if user.Avatar != "" {
			user.Avatar = "https://storage.yandexcloud.net/lms-user/" + user.Avatar
		}

		c.JSON(http.StatusOK, gin.H{
			"user": UserResponse{
				UUID:        user.ID,
				Login:       user.Login,
				Name:        user.Name,
				Surname:     user.Surname,
				Email:       user.Email,
				Phone:       user.Phone,
				PlaceWork:   user.PlaceWork,
				Position:    user.Position,
				CreatedDate: user.CreatedAt,
				URL:         user.Avatar,
			},
			"avatar": base64.StdEncoding.EncodeToString(buf),
		})
	}
}

func getAllUser(a app.IAppUser) gin.HandlerFunc {
	return func(c *gin.Context) {
		users, err := a.GetUsers(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"method": c.FullPath(),
				"err":    errors.Wrap(err, "get user domain").Error(),
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"users": users,
		})
	}
}
