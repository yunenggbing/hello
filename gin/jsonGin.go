package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
	"net/http"
)

/*
@Time : 2024/7/26 11:31
@Author : echo
@File : jsonGin
@Software: GoLand
@Description:
*/
type Login struct {
	User    string `form:"username" json:"user" uri:"user" xml:"user" binding:"required"`
	Pssword string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

func main() {
	r := gin.Default()

	r.POST("loginJSON", func(c *gin.Context) {
		var json Login
		//将request的body中的数据，自动按照json格式解析到结构体
		if err := c.ShouldBindBodyWithJSON(&json); err != nil {

			//返回错误信息
			// gin.H封装了map[string]interface{}
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if json.User != "root" || json.Pssword != "admin" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "200"})
	})

	r.POST("loginForm", func(c *gin.Context) {
		var form Login
		//Bind()默认解析并绑定form数据，根据Content-Type自动推断
		if err := c.Bind(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if form.User != "root" || form.Pssword != "admin" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "304"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "200"})
	})

	//Uri参数
	r.GET("/:user/:password", func(c *gin.Context) {
		var login Login
		if err := c.ShouldBindUri(&login); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if login.User != "root" || login.Pssword != "admin" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "200"})
	})

	//响应
	//json
	r.GET("/someJSON", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hey, this is json",
		})
	})
	//结构体
	r.GET("/someStruct", func(c *gin.Context) {
		var msg struct {
			Name    string
			Message string
			Number  int
		}
		msg.Name = "root"
		msg.Message = "hey, this is struct"
		msg.Number = 123
		c.JSON(http.StatusOK, msg)
	})
	//xml
	r.GET("/someXML", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{"message": "hey, this is xml"})
	})
	//YAML
	r.GET("/someYAML", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{"message": "hey, this is yaml"})
	})
	//protobuf格式,谷歌开发的高效存储读取的工具
	// 数组？切片？如果自己构建一个传输格式，应该是什么格式？
	r.GET("/someProtoBuf", func(c *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		label := "label"
		data := &protoexample.Test{
			Label: &label,
			Reps:  reps,
		}
		c.ProtoBuf(http.StatusOK, data)
	})
	//重定向
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
	})
	r.Run(":8000")
}
