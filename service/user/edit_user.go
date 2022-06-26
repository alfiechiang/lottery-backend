package user

import (
	"lottery/model"
	"lottery/serializer"

	"github.com/gin-gonic/gin"
)

type EditUserRequest struct {
	UserID int `uri:"userid" form:"userid"`
}

type EditUserResponse struct {
	ID           uint   `gorm:"primarykey"`
	UserName     string `gorm:"column:username"`
	Nickname     string
	Distributors string
	Role         string
	Status       int
	Avatar       string `gorm:"size:1000"`
	CreatedAt    model.DateTime
	UpdatedAt    model.DateTime
	DeletedAt    model.DateTime
}

func (user *EditUserRequest) EditUser(c *gin.Context) serializer.Response {

	res := EditUserResponse{}

	model.DB.Table("users").First(&res, user.UserID)

	return serializer.Response{
		Code: 200,
		Msg:  "編輯用戶成功",
		Data: res,
	}
}
