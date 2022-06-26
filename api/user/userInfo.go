package user

import (
	api "lottery/api"
	service "lottery/service/user"

	"github.com/gin-gonic/gin"
)

func UserInfo(c *gin.Context) {

	var service service.UserInfoRequest
	if err := c.ShouldBind(&service); err == nil {
		res := service.UserInfo(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, api.ErrorResponse(err))
	}

}
