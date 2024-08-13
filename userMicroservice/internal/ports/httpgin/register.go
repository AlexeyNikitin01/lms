package httpgin

import (
	"github.com/gin-gonic/gin"

	"lms-user/internal/app"
)

func register(a app.IAppUser) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request UserRequest
		if err := c.Bind(&request); err != nil {
			return
		}

		user, err := a.Register(c, request.Login, request.Password, request.Email)
		if err != nil {
			return
		}

		c.JSON(200, gin.H{
			"user": UserResponse{
				Uuid:        user.ID,
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

		return
	}
}
