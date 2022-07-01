package six_lottery

import (
	"lottery/model"
	"lottery/serializer"

	"github.com/gin-gonic/gin"
)

type SixlotteryListRequest struct{}

func (request *SixlotteryListRequest) SixlotteryList(c *gin.Context) serializer.Response {

	var list []model.SixLotteryList
	model.DB.Table("six_lottery_list").Find(&list)

	return serializer.Response{
		Code: 200,
		Data: list,
		Msg:  "列表成功",
	}
}
