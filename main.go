package main

import (
	"go_scaffold_project/config"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Run(config.Port)

}
