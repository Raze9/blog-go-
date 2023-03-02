package routes

import (
	"GOproject/blog2/controller"
	"GOproject/blog2/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	// 允许跨域访问
	// 注册
	r := gin.Default()
	v1 := r.Group("/api")
	{
		v1.POST("/register", controller.Register)
		// 登录
		v1.POST("/login", controller.Login)
		// 登录获取用户信息
		v1.GET("/user", middleware.AuthMiddleware(), controller.GetInfo)
	}
	return r
}
