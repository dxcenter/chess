package main

import (
	"github.com/gin-gonic/gin"
	m "github.com/dxcenter/chess/serverMethods"
)

func setupRouter(r *gin.Engine) {
	r.Static("/frontend", "frontend")
	r.StaticFile("/", "frontend/index.html")

	r.GET("/ping.json", m.Ping)
	r.GET("/auth.json", m.Auth)
}
