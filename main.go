package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

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
