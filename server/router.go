package server

import (
	"lottery/api/menu"
	"lottery/api/six_lottery"
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

			v1.GET("menu", menu.MenuList)
			v1.GET("sixlottery/colly", six_lottery.Colly)
			v1.GET("sixlottery", six_lottery.SixlotteryList)

		}

	}
	return r
}
