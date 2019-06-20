package main

import (
	"./utils"
	"github.com/gin-gonic/gin"
)

func main() {

	driver := gin.Default()

	driver.GET("/info/:name", utils.InfoHandler)

	driver.Run()

}
