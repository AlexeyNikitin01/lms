package httpgin

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"lms-user/internal/app"
)

func Server(addr string, app app.IAppUser) *http.Server {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(auth(app))

	s := &http.Server{
		Addr:    addr,
		Handler: router,
	}
	AppRouter(router.Group("user"), app)
	return s
}
