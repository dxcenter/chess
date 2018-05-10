package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	setupRouter(r)
	r.Run()
}
