package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

/*
@Time : 2024/7/25 9:51
@Author : echo
@File : firstGin
@Software: GoLand
@Description:
*/
func main() {
	//1.创建一个默认的gin引擎
	r := gin.Default()
	//绑定路由规则，执行的函数。 gin.Context 封装了request和response
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, Gin!")
	})
	r.POST("/helloPost", func(c *gin.Context) {
		c.String(http.StatusBadGateway, "Hello, Gin Post!")
	})
	//get 获取路由参数
	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		action = strings.Trim(action, "/")
		c.String(http.StatusOK, name+" is "+action)
	})
	r.GET("/user1", func(c *gin.Context) {
		name := c.DefaultQuery("username", "echo")
		c.String(http.StatusOK, fmt.Sprintf("Hello %s", name))
	})
	//post 表单
	r.POST("/form", func(c *gin.Context) {
		types := c.DefaultPostForm("type", "post")
		username := c.PostForm("username")
		password := c.PostForm("password")
		c.String(http.StatusOK, fmt.Sprintf("username:%s,password:%s,type:%s", username, password, types))
	})
	//单个文件上传 保存到项目的目录中
	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.String(500, "上传文件失败")
		}
		c.SaveUploadedFile(file, file.Filename)
		c.String(http.StatusOK, "上传成功："+file.Filename)
	})
	//多个文件上传
	r.MaxMultipartMemory = 8 << 20 // 8 MiB
	r.POST("/uploads", func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		}
		files := form.File["files"]
		for _, file := range files {
			if err := c.SaveUploadedFile(file, file.Filename); err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
				return
			}
		}
		c.String(http.StatusOK, fmt.Sprintf("upload ok %d files", len(files)))
	})
	//启动服务
	if err := r.Run(":8000"); err != nil {
		panic("启动失败" + err.Error())
	}

}
