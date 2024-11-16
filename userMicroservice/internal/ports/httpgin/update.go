package httpgin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"github.com/lms-user/internal/app"
	"github.com/lms-user/internal/repository/pg/entity"
)

func updateUser(a app.IAppUser) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req UpdateUserRequest
		if err := c.Bind(&req); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"method": c.FullPath(),
				"err":    errors.Wrap(err, "user update request").Error(),
			})

			return
		}

		currentUser := FromContext(c)

		if !checkUser(currentUser, req.UUID) {
			c.JSON(http.StatusInternalServerError, gin.H{
				"method": c.FullPath(),
				"error":  errors.New("you don`t update user").Error(),
			})

			return
		}

		userForUpdate := entity.User{
			ID:        req.UUID,
			Login:     req.Login,
			Name:      req.Name,
			Surname:   req.Surname,
			Phone:     req.Phone,
			Position:  req.Position,
			PlaceWork: req.PlaceWork,
			Email:     req.Email,
		}

		user, err := a.UpdateUser(c, &userForUpdate)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"err":    errors.Wrap(err, "gin: failed to update user"),
				"method": c.FullPath(),
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"user": &UserResponse{
				UUID:        user.ID,
				Login:       user.Login,
				Name:        user.Name,
				Surname:     user.Surname,
				Email:       user.Email,
				Phone:       user.Phone,
				PlaceWork:   user.PlaceWork,
				Position:    user.Position,
				CreatedDate: user.CreatedAt,
			},
		})
	}
}

func checkUser(user *entity.User, userUUID string) bool {
	return user.Role == "admin" || user.ID == userUUID
}
