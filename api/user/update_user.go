package user

import (
	"lottery/api"
	service "lottery/service/user"

	"github.com/gin-gonic/gin"
)

func UpdateUser(c *gin.Context) {
	var service service.UpdateUserRequest
	if err := c.ShouldBind(&service); err == nil {
		res := service.UpdateUser(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, api.ErrorResponse(err))
	}
}
