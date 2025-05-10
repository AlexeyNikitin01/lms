package httpgin

import (
	"manufactures/internal/app"

	"github.com/gin-gonic/gin"
)

func AppRouter(r *gin.RouterGroup, a app.AppManfs) {
	r.GET("ping", ping(a))
	r.GET("data", data(a))
	r.GET("vacancies/:company", vacancies(a))
}
