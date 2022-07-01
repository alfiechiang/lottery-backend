package six_lottery

import (
	"lottery/api"
	service "lottery/service/six_lottery"
	"github.com/gin-gonic/gin"
)

func Colly(c *gin.Context) {

	var service service.CollyRequest
	if err := c.ShouldBind(&service); err == nil {
		res := service.Colly(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, api.ErrorResponse(err))
	}
	

}