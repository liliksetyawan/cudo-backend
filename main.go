package main

import (
	"cudo/test/config"
	"cudo/test/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	router := gin.Default()
	routers.UserRoutes(router)
	routers.TransactionRouters(router)
	router.Run(":8080")
}
