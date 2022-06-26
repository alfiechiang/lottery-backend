package user

import (
	"lottery/model"
	"lottery/serializer"
	"lottery/util/hash"

	"github.com/gin-gonic/gin"
)

type UpdateUserRequest struct {
	UserName     string `json:"username" gorm:"column:username"`
	Password     string
	Nickname     string
	Distributors string
	Role         string
	Status       int
	Avatar       string
}

func (request *UpdateUserRequest) UpdateUser(c *gin.Context) serializer.Response {

	userid := c.Param("userid")
	var user model.User

	model.DB.Table("users").First(&user, userid)
	user.UserName = request.UserName
	user.Password = hash.HashMake(request.Password)
	user.Nickname = request.Nickname
	user.Distributors = request.Distributors
	user.Role = request.Role
	user.Status = request.Status
	user.Avatar = request.Avatar
	if err := model.DB.Save(&user).Error; err != nil {
		return serializer.Err(10004)
	}
	return serializer.Response{
		Code: 200,
		Msg:  "修改用戶成功",
		Data: make([]string, 0),
	}

}
