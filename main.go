package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type ReturnData struct {
	Status  string   `json:"status"`
	Message string   `json:"message"`
	Data    UserInfo `json:"data"`
}
type UserInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func getUserInfo(c *gin.Context) {
	username := c.Query("username")
	password, ok := c.GetQuery("password")
	data := ReturnData{
		Status:  "200",
		Message: "ok",
		Data: UserInfo{
			Username: username,
			Password: password,
		},
	}
	if ok {
		c.JSON(200, data)
	}
}

func getUserInfo_formData(c *gin.Context) {
	username := c.PostForm("username")
	password, _ := c.GetPostForm("password")
	data := ReturnData{
		Status:  "200",
		Message: "ok",
		Data: UserInfo{
			Username: username,
			Password: password,
		},
	}
	c.JSON(200, data)
}

func getUserInfo_rawData(c *gin.Context) {
	jsondata := ReturnData{}
	if err := c.ShouldBindJSON(&jsondata); err != nil {
		data := ReturnData{
			Status:  "400",
			Message: "绑定错误"}
		c.JSON(200, data)
	} else {
		data := ReturnData{
			Status:  "200",
			Message: "ok",
			Data: UserInfo{
				Username: jsondata.Data.Username,
				Password: jsondata.Data.Password,
			},
		}
		c.JSON(200, data)
	}
	fmt.Println(jsondata)
}

func golableMiddleWare(c *gin.Context) {
	_, ok := c.GetQuery("token")
	if !ok {
		c.String(200, "get token error")
		c.Abort()
	}
	fmt.Println("before abort")
}

func partMiddleWare(c *gin.Context) {
	username, _ := c.GetQuery("username")
	if username != "admin" {
		c.String(200, "user is not admin!")
		c.Abort()
	}

}
func main() {
	r := gin.Default()
	r.GET("/", getUserInfo)
	apiGroup := r.Group("/api")
	apiGroup.Use(golableMiddleWare)
	userGroup := apiGroup.Group("/user")
	userGroup.GET("/login", partMiddleWare, getUserInfo_rawData)
	userGroup.GET("/logout")
	// productGroup := apiGroup.Group("/product")
	r.Run()

}
