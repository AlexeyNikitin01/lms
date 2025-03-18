package httpgin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"course/internal/app"
)

func Server(addr string, app app.ICourseApp) *http.Server {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	router.Use(corsMiddleware())
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(prometheusMiddleware())

	s := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	router.GET("metrics", gin.WrapH(promhttp.Handler()))
	AppRouter(router.Group("course"), app)

	return s
}
