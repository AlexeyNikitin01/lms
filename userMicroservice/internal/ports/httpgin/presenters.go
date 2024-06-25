package httpgin

import (
	"github.com/gin-gonic/gin"
)

type User struct {
}

type Error struct {
}

func ResponseUser() *gin.H {
	return &gin.H{}
}
