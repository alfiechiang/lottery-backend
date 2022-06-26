package user

import (
	"lottery/api"
	service "lottery/service/user"

	"github.com/gin-gonic/gin"
)

func UserList(c *gin.Context) {
	var service service.UserListRequest
	if err := c.ShouldBind(&service); err == nil {
		res := service.UserList(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, api.ErrorResponse(err))
	}
}
