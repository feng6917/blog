package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	//定义路由的GET方法及响应处理函数
	r.GET("/hello", func(c *gin.Context) {
		//将发送的信息封装成JSON发送给浏览器
		c.JSON(http.StatusOK, gin.H{
			//这是我们定义的数据
			"message": "快速入门",
		})
	})
	r.Run() //默认在本地8080端口启动服务
}
