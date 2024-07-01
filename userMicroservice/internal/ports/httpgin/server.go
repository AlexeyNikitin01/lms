package httpgin

import (
	"lms-user/internal/app"

	"github.com/gin-gonic/gin"
	"net/http"
)

func Server(addr string, app app.IAppUser) *http.Server {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	s := &http.Server{
		Addr:    addr,
		Handler: router,
	}
	AppRouter(router.Group("user"), app)
	return s
}
