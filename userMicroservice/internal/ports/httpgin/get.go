package httpgin

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"lms-user/internal/app"
)

func getUser(a app.IAppUser) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req GetUserReq
		if err := c.Bind(&req); err != nil {
			c.JSON(500, gin.H{
				"err":    errors.Wrap(err, "get user request"),
				"method": c.FullPath(),
			})
			return
		}
		fmt.Println(req.UUID)

		user, err := a.GetUser(c, req.UUID)
		if err != nil {
			c.JSON(500, gin.H{
				"err":    errors.Wrap(err, "get user domain"),
				"method": c.FullPath(),
			})
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
	}
}
