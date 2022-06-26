package user

import (
	"lottery/api"
	service "lottery/service/user"

	"github.com/gin-gonic/gin"
)

func EditUser(c *gin.Context) {
	var service service.EditUserRequest
	if err := c.ShouldBindUri(&service); err == nil {
		res := service.EditUser(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, api.ErrorResponse(err))
	}
}
