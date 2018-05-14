package main

import (
	cfg "github.com/dxcenter/chess/config"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg.Reload()

	r := gin.Default()
	setupRouter(r)
	r.Run()
}
