package main

import (
	m "github.com/dxcenter/chess/serverMethods"
	"github.com/gin-gonic/gin"
)

func setupRouter(r *gin.Engine) {
	r.Static("/frontend", "frontend")
	r.StaticFile("/", "frontend/index.html")

	r.GET("/ping.json", m.Ping)

	// JWT
	jwt := newJwtMiddleware()
	r.GET("/auth.json", jwt.LoginHandler)
	r.GET("/refresh_token.json", jwt.RefreshHandler)
}
