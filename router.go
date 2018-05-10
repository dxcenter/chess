package main

import (
	m "github.com/dxcenter/chess/serverMethods"
	"github.com/gin-gonic/gin"
)

func setupRouter(r *gin.Engine) {
	r.Static("/frontend", "frontend/build")
	r.Static("/static", "frontend/build/static")
	r.StaticFile("/", "frontend/build/index.html")
	r.StaticFile("/login", "frontend/build/index.html")

	for _, file := range []string{"index.html", "service-worker.js"} {
		r.StaticFile(file, "frontend/build/"+file)
	}

	r.GET("/ping.json", m.Ping)

	// JWT
	jwt := newJwtMiddleware()
	r.POST("/auth.json", jwt.LoginHandler)
	r.POST("/refresh_token.json", jwt.RefreshHandler)
}
