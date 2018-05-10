package main

import (
	m "github.com/dxcenter/chess/serverMethods"
	"github.com/gin-gonic/gin"
)

func setupRouter(r *gin.Engine) {
	devMode := true

	// My methods
	r.GET("/ping.json",        m.Ping)
	r.GET("/game_status.json", m.GameStatus)
	r.POST("/new_game.json",   m.NewGame)
	r.POST("/move.json",       m.Move)

	// Frontend
	if !devMode {
		r.Static("/frontend", "frontend/build")
		r.Static("/static", "frontend/build/static")
		r.Static("/css", "frontend/build/css")
		r.StaticFile("/", "frontend/build/index.html")
		r.StaticFile("/login", "frontend/build/index.html")
		for _, file := range []string{"index.html", "service-worker.js"} {
			r.StaticFile(file, "frontend/build/"+file)
		}
	}

	// node.js
	r.GET("/sockjs-node/:id1/:id2/websocket", m.ProxyToNodejs)
	if devMode {
		r.GET("/", m.ProxyToNodejs)
		r.("/frontend", "frontend/build")
		r.("/static", "frontend/build/static")
		r.("/css", "frontend/build/css")
		r.GET("/", m.ProxyToNodejs)
		r.GET("/login", m.ProxyToNodejs)
		for _, file := range []string{"index.html", "service-worker.js"} {
			r.GET("/"+file, m.ProxyToNodejs)
		}
	}

	// JWT
	jwt := newJwtMiddleware()
	r.POST("/auth.json", jwt.LoginHandler)
	r.POST("/refresh_token.json", jwt.RefreshHandler)
}
