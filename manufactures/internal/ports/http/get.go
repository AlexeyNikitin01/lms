package httpgin

import (
	"github.com/gin-gonic/gin"

	"github.com/ekomobile/dadata/v2/api/suggest"

	"manufactures/internal/app"
)

func data(a app.AppManfs) gin.HandlerFunc {
	return func(c *gin.Context) {
		params := suggest.PartyByIDParams{
			Query:      "7610052644",
			BranchType: ToPtr(suggest.PartyBranchTypeMain),
		}

		suggestions, err := a.D.Api.PartyByID(c, &params)
		if err != nil {
			return
		}

		c.JSON(200, gin.H{
			"data": suggestions,
		})
	}
}

func ToPtr[T any](x T) *T {
	return &x
}
