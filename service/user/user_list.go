package user

import (
	"lottery/model"
	"lottery/serializer"

	"github.com/gin-gonic/gin"
)

type UserListRequest struct {
}

func (request *UserListRequest) UserList(c *gin.Context) serializer.Response {

	var list []model.User
	model.DB.Table("users").Find(&list)

	return serializer.Response{
		Code: 200,
		Msg:  "列表成功",
		Data: list,
	}
}
