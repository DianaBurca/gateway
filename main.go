package main

import (
	"github.com/DianaBurca/gateway/utils"
	"github.com/gin-gonic/gin"
)

func main() {

	driver := gin.Default()

	driver.GET("/info", utils.InfoHandler)
	driver.GET("/.well-known/live", utils.Health)
	driver.GET("/.well-known/ready", utils.Health)

	driver.Run()

}
