package six_lottery

import (
	"lottery/api"
	service "lottery/service/six_lottery"

	"github.com/gin-gonic/gin"
)

func SixlotteryList(c *gin.Context) {

	var service service.SixlotteryListRequest
	if err := c.ShouldBind(&service); err == nil {
		res := service.SixlotteryList(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, api.ErrorResponse(err))
	}

}
