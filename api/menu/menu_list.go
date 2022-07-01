package menu

import (
	"lottery/api"
	service "lottery/service/menu"

	"github.com/gin-gonic/gin"
)

func MenuList(c *gin.Context) {

	var service service.MenuListRequest
	if err := c.ShouldBind(&service); err == nil {
		res := service.MenuList(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, api.ErrorResponse(err))
	}

}
