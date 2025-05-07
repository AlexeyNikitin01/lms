package httpgin

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"manufactures/internal/app"
	user "manufactures/pkg/grpc"
)

func Server(addr string, app app.AppManfs, userClient user.UserServiceClient) *http.Server {
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

	group := router.Group("map")
	AppRouter(group, app)

	return s
}
