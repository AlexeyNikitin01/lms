package httpgin

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"lms-user/internal/app"
)

func authByLoginPassword(a app.IAppUser) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req AuthByLogPassRequest
		if err := c.Bind(&req); err != nil {
			c.JSON(500, gin.H{
				"method": c.FullPath(),
				"err":    errors.Wrap(err, "auth by login password gin").Error(),
			})
			return
		}

		_, token, err := a.AuthByLoginPassword(c, req.Login, req.Password)
		if err != nil {
			c.JSON(500, gin.H{
				"method": c.FullPath(),
				"err":    errors.Wrap(err, "auth by login password gin").Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"data": AuthByLogPassResponse{
				Token:   token.Token,
				Refresh: token.Refresh,
			},
		})

		return
	}
}
