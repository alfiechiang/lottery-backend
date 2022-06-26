package user

import (
	"lottery/model"
	"lottery/serializer"

	"github.com/gin-gonic/gin"
)

type DeleteUserRequest struct {
	UserID int `uri:"userid" form:"userid"`
}

func (user *DeleteUserRequest) DeleteUser(c *gin.Context) serializer.Response {

	if err := model.DB.Delete(&model.User{}, user.UserID).Error; err != nil {
		return serializer.Err(10003)
	}

	return serializer.Response{
		Code: 200,
		Msg:  "刪除用戶成功",
		Data: make([]string, 0),
	}
}
