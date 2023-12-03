package main

import (
	"fmt"
	"go_scaffold_project/config"
	"go_scaffold_project/middlewares"
	"go_scaffold_project/routers"
	"go_scaffold_project/utils/jwts"
	"go_scaffold_project/utils/logs"

	"github.com/gin-gonic/gin"
)

func main() {
	logs.Info(nil, "test")
	fmt.Println(config.JwtSecret)

	token, _ := jwts.GenToken("admin")
	clam, err := jwts.ParseToken(token)
	fmt.Println(clam, err)

	r := gin.Default()
	routers.RegisterRouters(r)
	r.Use(middlewares.JwtCheck)
	r.Run()
}
