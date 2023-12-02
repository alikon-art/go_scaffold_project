package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// "go_scaffold_project/config"

// "github.com/gin-gonic/gin"

type MyCustomClaims struct {
	UserID     int
	Username   string
	GrantScope string
	jwt.RegisteredClaims
}

func main() {
	// r := gin.Default()
	// r.Run(config.Port)
	jwttoken := "test"
	jj := []byte(jwttoken)
	fmt.Println(jj)

	claim := MyCustomClaims{
		UserID:     000001,
		Username:   "Tom",
		GrantScope: "read_user_info",
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "Auth_Server",                                   // 签发者
			Subject:   "Tom",                                           // 签发对象
			Audience:  jwt.ClaimStrings{"Android_APP", "IOS_APP"},      //签发受众
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),   //过期时间
			NotBefore: jwt.NewNumericDate(time.Now().Add(time.Second)), //最早使用时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                  //签发时间
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	fmt.Println(token)
	ss, err := token.SignedString(jj)
	fmt.Println(ss, err)

	tokenss, err := jwt.ParseWithClaims(ss, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("AllYourBase"), nil
	})
	print(tokenss)

}
