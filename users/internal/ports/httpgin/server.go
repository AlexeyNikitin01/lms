package httpgin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"lms-user/internal/app"
)

func Server(addr string, app app.IAppUser) *http.Server {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	router.Use(corsMiddleware())
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(auth(app))

	s := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	userGroup := router.Group("user")
	AppRouter(userGroup, app)

	metricsGroup := router.Group("metrics")
	metricsGroup.GET("", gin.WrapH(promhttp.Handler()))

	return s
}
