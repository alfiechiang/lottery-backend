package user

import (
	"lottery/model"
	"lottery/serializer"

	"github.com/gin-gonic/gin"
)

type UserInfoRequest struct {
	AuthUser    model.User `json:"auth_user"  form:"auth_user"`
	AccessToken string
}

type UserInforesponse struct {
	Roles    []string `json:"roles"`
	UserName string   `json:"username"`
}

func (info *UserInfoRequest) UserInfo(c *gin.Context) serializer.Response {

	user := c.MustGet("auth_user").(model.User)

	res := UserInforesponse{
		Roles:    []string{"admin"},
		UserName: user.UserName,
	}

	return serializer.Response{
		Code: 200,
		Msg:  "成功",
		Data: res,
	}
}
