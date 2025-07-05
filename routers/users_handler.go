package routers

import (
	"cudo/test/service"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	router.POST("/users", service.CreateUser)
}
