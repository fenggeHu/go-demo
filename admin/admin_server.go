package admin

import (
	"github.com/gin-gonic/gin"
	"jinfeng.hu/goapi/admin/controller"
)

// Server 后台服务 - 分组
func Server() {
	router := gin.Default()
	// 设置一个get请求的路由，url为/ping, 处理函数（或者叫控制器函数）是一个闭包函数。
	router.GET("/ping", func(c *gin.Context) {
		// 通过请求上下文对象Context, 直接往客户端返回一个json
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	user := router.Group("/user/v1")
	{
		// user login
		user.POST("/login", controller.Login)
	}

	router.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
