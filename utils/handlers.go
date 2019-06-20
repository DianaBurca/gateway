package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InfoHandler(c *gin.Context) {

	driver := gin.Default()
	name := c.Param("name")
	c.String(http.StatusOK, "Hello %s", name)
	v1 := driver.Group("/v1")
	{
		v1.GET("/store/" + name) //cron(C)
		v1.GET("/read/" + name)  //info-reader(A)

	}

}
