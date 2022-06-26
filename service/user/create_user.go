package user

import (
	"lottery/model"
	"lottery/serializer"

	"github.com/gin-gonic/gin"
)

type CreateUserRequest struct {
	UserName     string `json:"username" gorm:"column:username"`
	Password     string
	Nickname     string
	Distributors string
	Role         string
	Status       int
	Avatar       string
}

func (user *CreateUserRequest) CreateUser(c *gin.Context) serializer.Response {

	if err := model.DB.Table("users").Create(&user).Error; err != nil {
		return serializer.Err(10002)
	}

	return serializer.Response{
		Code: 200,
		Msg:  "創建用戶成功",
		Data: make([]string, 0),
	}

}
