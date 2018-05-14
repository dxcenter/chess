package main

import (
	m "github.com/dxcenter/chess/serverMethods"
	"github.com/gin-gonic/gin"
)

func setupRouter(r *gin.Engine) {
	// My methods
	r.GET("/ping.json",        m.Ping)
	r.GET("/game_status.json", m.GameStatus)
	r.POST("/new_game.json",   m.NewGame)
	r.POST("/move.json",       m.Move)

	// Frontend
	r.Static("/frontend", "frontend/build")
	r.Static("/static", "frontend/build/static")
	r.Static("/css", "frontend/build/css")
	r.StaticFile("/", "frontend/build/index.html")
	r.StaticFile("/login", "frontend/build/index.html")
	for _, file := range []string{"index.html", "service-worker.js"} {
		r.StaticFile(file, "frontend/build/"+file)
	}

	// JWT
	jwt := newJwtMiddleware()
	r.POST("/auth.json", jwt.LoginHandler)
	r.POST("/refresh_token.json", jwt.RefreshHandler)
}
