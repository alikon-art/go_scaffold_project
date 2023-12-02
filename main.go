package main

import (
	"fmt"
	"go_scaffold_project/config"
	"go_scaffold_project/utils/jwts"
	"go_scaffold_project/utils/logs"
)

func main() {
	logs.Info(nil, "test")
	fmt.Println(config.JwtSecret)

	token, _ := jwts.GenToken("admin")
	clam, err := jwts.ParseToken(token)
	fmt.Println(clam, err)
}
