package user

import (
	"lottery/auth"
	"lottery/model"
	"lottery/serializer"
	"lottery/util/hash"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	UserName string
	Password string
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

func (login *LoginRequest) Login(c *gin.Context) serializer.Response {

	var user model.User
	model.DB.Table("users").Where("username", login.UserName).Find(&user)

	if user == (model.User{}) {
		return serializer.Err(10000)
	}

	if check := hash.Check(user.Password, login.Password); check != nil {
		return serializer.Err(10001)
	}

	token, _ := auth.GenToken(user)

	return serializer.Response{
		Code: 200,
		Msg:  "登入成功",
		Data: LoginResponse{AccessToken: token},
	}
}
