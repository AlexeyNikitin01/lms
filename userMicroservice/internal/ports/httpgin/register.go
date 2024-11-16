package httpgin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"github.com/AlexeyNikitin01/lms-user/internal/app"
)

func register(a app.IAppUser) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request UserRequest
		if err := c.Bind(&request); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"method": c.FullPath(),
				"err":    errors.Wrap(err, "register request").Error(),
			})

			return
		}

		user, err := a.Register(c, request.Login, request.Password, request.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"method": c.FullPath(),
				"err":    errors.Wrap(err, "register domain").Error(),
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"user": UserRegisterResponse{
				UUID:  user.ID,
				Login: user.Login,
				Email: user.Email,
			},
		})
	}
}
