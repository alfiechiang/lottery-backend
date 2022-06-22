package auth

import (
	"errors"
	"lottery/model"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type MyClaims struct {
	AuthUser model.User
	jwt.StandardClaims
}

type MyCustClaims struct {
	CustData interface{}
	jwt.StandardClaims
}

const TokenExpireDuration = time.Hour * 2

var MySecret = []byte("d8YDwSXtWTUrKSKP") // []byte(os.Getenv("Token_Pwd"))

func GenToken(auth model.User) (string, error) {

	// 创建一个我们自己的声明
	c := MyClaims{
		auth,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "my-project",                               // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(MySecret)
}

func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

// 自定 custData 用 interface，能接受任何struct，但還原的時候比較麻煩
// 需要 json.Marshal json.Unmarshal 兩個方式轉成正確的 struct
// 第2個參數 持續幾秒
var TGSecret = []byte("vXqrHRmAaDYNgTWp") // []byte(os.Getenv("Token_Pwd"))

func GenCustToken(custData interface{}, args ...int) (string, error) {
	sec := 30
	if len(args) > 0 {
		sec = args[0]
	}

	// 创建一个我们自己的声明
	c := MyCustClaims{
		custData,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * time.Duration(sec)).Unix(), // 过期时间
			Issuer:    "backend_api_cust_token",                                // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(TGSecret)
}

func ParseCustToken(tokenString string) (*MyCustClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &MyCustClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return TGSecret, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*MyCustClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
