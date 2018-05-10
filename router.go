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
	if devMode {
		r.POST("/sockjs-node/*rest", m.ProxyToNodejs)
		r.GET("/sockjs-node/*rest", m.ProxyToNodejs)
		r.GET("/__webpack_dev_server__/*rest", m.ProxyToNodejs)

		r.GET("/frontend/*rest", m.ProxyToNodejs)
		r.GET("/static/*rest", m.ProxyToNodejs)
		r.GET("/css/*rest", m.ProxyToNodejs)
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
