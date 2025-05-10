package httpgin

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"lab/internal/app"
	user "lab/pkg/grpc"
)

func Server(addr string, app *app.Lab, userClient user.UserServiceClient) *http.Server {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	router.Use(corsMiddleware())
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(auth(userClient))

	s := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	group := router.Group("labs")
	AppRouter(group, app)

	return s
}
