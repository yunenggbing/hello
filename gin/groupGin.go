package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

/*
@Time : 2024/7/26 10:46
@Author : echo
@File : groupGin
@Software: GoLand
@Description:
*/
func main() {
	r := gin.Default()
	v1 := r.Group("/v1")

	{
		v1.GET("/login", login)
		v1.GET("/submit", submit)
	}

	v2 := r.Group("/v2")
	{

		v2.POST("/login", login)
		v2.POST("/submit", submit)
	}

	r.Run(":8000")
}

func login(c *gin.Context) {
	name := c.DefaultQuery("name", "anonymous")
	c.String(200, fmt.Sprintf("hello %s", name))
}
func submit(c *gin.Context) {
	name := c.DefaultQuery("name", "gaga")
	c.String(200, fmt.Sprintf("hello %s", name))
}
