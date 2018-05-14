package main

import (
	cfg "github.com/dxcenter/chess/config"
	jwt "github.com/dxcenter/chess/jwt"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg.Reload()
	jwt.InitJwtMiddleware()

	r := gin.Default()
	setupRouter(r)
	r.Run()
}
