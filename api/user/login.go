package user

import (
	api "lottery/api"
	service "lottery/service/user"

	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {

	var service service.LoginRequest
	if err := c.ShouldBind(&service); err == nil {
		res := service.Login(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, api.ErrorResponse(err))
	}

}
