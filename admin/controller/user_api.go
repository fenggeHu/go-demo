package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"jinfeng.hu/goapi/repository/service"
	"log"
	"net/http"
)

type LoginDO struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	user := c.Query("username")
	pwd := c.Query("password")
	log.Printf("Query Params -> username: %s, password: %s", user, pwd)
	username := c.PostForm("username")
	password := c.PostForm("password")
	log.Printf("PostForm -> username: %s, password: %s", username, password)

	loginDO := LoginDO{}
	if err := c.ShouldBindBodyWith(&loginDO, binding.JSON); err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		return
	}
	log.Printf("loginDO: %s", loginDO)

	// 设置http响应头（设置Header） - 要放在响应请求之前
	c.Header("Author","jinfeng")
	c.Header("ORG","HZ")
	_, err := service.Login(loginDO.Username, loginDO.Password)
	if nil != err {
		// 以字符串方式响应请求
		c.String(http.StatusForbidden, "login failure: %s", err)
	} else {
		// 以json格式响应请求
		c.JSON(http.StatusOK, loginDO)
	}
	// 以xml格式响应请求雷同 c.XML(200, u)
	// 以文件格式响应请求 c.File("/var/www/1.jpg") // c.FileAttachment("/var/www/1.jpg", "1.jpg")
}
