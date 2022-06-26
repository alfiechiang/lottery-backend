package user

import (
	"lottery/api"
	service "lottery/service/user"

	"github.com/gin-gonic/gin"
)

func DeleteUser(c *gin.Context) {
	var service service.DeleteUserRequest
	if err := c.ShouldBindUri(&service); err == nil {
		res := service.DeleteUser(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, api.ErrorResponse(err))
	}
}
