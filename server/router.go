package server

import (
	"lottery/api/user"
	"lottery/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser())

	// 路由
	v1 := r.Group("/api/v1")
	{
		v1.POST("login", user.LoginHandler)

		v1.Use(middleware.JWTAuthMiddleware())
		{
			v1.POST("userinfo", user.UserInfo)

		}

	}
	return r
}
