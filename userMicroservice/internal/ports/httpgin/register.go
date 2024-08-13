package httpgin

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"lms-user/internal/app"
)

func register(a app.IAppUser) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request UserRequest
		if err := c.Bind(&request); err != nil {
			c.JSON(500, gin.H{
				"method": c.FullPath(),
				"err":    errors.Wrap(err, "register request").Error(),
			})
			return
		}

		user, err := a.Register(c, request.Login, request.Password, request.Email)
		if err != nil {
			c.JSON(500, gin.H{
				"method": c.FullPath(),
				"err":    errors.Wrap(err, "register domain").Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"user": UserRegisterResponse{
				Uuid:  user.ID,
				Login: user.Login,
				Email: user.Email,
			},
		})

		return
	}
}
