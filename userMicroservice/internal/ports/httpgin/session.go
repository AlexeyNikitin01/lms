package httpgin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"github.com/lms-user/internal/app"
)

func authByLoginPassword(a app.IAppUser) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req AuthByLogPassRequest
		if err := c.Bind(&req); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"method": c.FullPath(),
				"err":    errors.Wrap(err, "auth by login password gin").Error(),
			})

			return
		}

		user, token, err := a.AuthByLoginPassword(c, req.Login, req.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"method": c.FullPath(),
				"err":    errors.Wrap(err, "auth by login password gin").Error(),
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": AuthByLogPassResponse{
				Token:    token.Token,
				Refresh:  token.Refresh,
				UserUUID: user.ID,
			},
		})
	}
}
