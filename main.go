package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("leonzhangxf api start")

	engine := gin.Default()

	engine.GET("/", func(context *gin.Context) {
		context.String(200, "%s", "leon")
	})

	_ = engine.Run()
}
