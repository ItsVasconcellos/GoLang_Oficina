package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
