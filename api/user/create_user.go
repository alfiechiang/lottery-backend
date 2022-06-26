package user

import (
	"lottery/api"
	service "lottery/service/user"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var service service.CreateUserRequest
	if err := c.ShouldBind(&service); err == nil {
		res := service.CreateUser(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, api.ErrorResponse(err))
	}
}
