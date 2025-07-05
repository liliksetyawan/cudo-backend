package routers

import (
	"cudo/test/service"
	"github.com/gin-gonic/gin"
)

func TransactionRouters(router *gin.Engine) {
	router.GET("/api/v1/fraud-detection", service.GetFraudTransaction)
}
