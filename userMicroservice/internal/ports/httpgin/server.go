package httpgin

import (
	"edu-material/userMicroservice/internal/app"

	"net/http"
	"github.com/gin-gonic/gin"
)

func Server(addr string, app app.AppUser) *http.Server {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	s := &http.Server{
		Addr: addr,
		Handler: router,
	}
	AppRouter(router.Group("user"), app) 
	return s
}
