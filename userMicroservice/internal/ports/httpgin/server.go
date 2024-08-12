package httpgin

import (
	"net/http"

	"lms-user/internal/app"

	"github.com/gin-gonic/gin"
)

func Server(addr string, app app.IAppUser) *http.Server {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	s := &http.Server{
		Addr:    addr,
		Handler: router,
	}
	AppRouter(router.Group("user"), app)
	return s
}
