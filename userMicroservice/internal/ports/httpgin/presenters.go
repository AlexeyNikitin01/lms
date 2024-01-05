package httpgin

import (
	"edu-material/userMicroservice/internal/user"

	"github.com/gin-gonic/gin"
)

type User struct {

}

type Error struct {

}

func ResponseUser(u *user.User) *gin.H {
	return &gin.H{

	}
}
