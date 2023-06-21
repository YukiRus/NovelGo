package main

import (
	"fmt"
	"novelgo/api/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println()

	r := gin.Default()

	routers.RegisterRouter(r)

	r.Run(":58000")
}
