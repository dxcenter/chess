package main

import (
	jwt "github.com/dxcenter/chess/jwt"
	m "github.com/dxcenter/chess/serverMethods"
	mw "github.com/dxcenter/chess/serverMethods/middleware"
	"github.com/gin-gonic/gin"
)

func setupRouter(r *gin.Engine) {
	// JWT
	jwtMiddleware := jwt.GetJwtMiddleware()
	authed := r.Group("/")
	authed.Use(jwtMiddleware.MiddlewareFunc()) // require to be authed
	authed.Use(mw.Authed)                      // some routines for an already authed
	r.POST("/auth.json", jwtMiddleware.LoginHandler)
	r.POST("/refresh_token.json", jwtMiddleware.RefreshHandler)

	// My methods
	r.GET("/ping.json", m.Ping)
	authed.GET("/whoami.json", m.Whoami)
	authed.GET("/game_status.json", m.GameStatus)
	authed.POST("/new_game.json", m.NewGame)
	authed.POST("/move.json", m.Move)

	// Frontend
	r.Static("/frontend", "frontend/build")
	r.Static("/static", "frontend/build/static")
	r.Static("/css", "frontend/build/css")
	r.StaticFile("/", "frontend/build/index.html")
	r.StaticFile("/login", "frontend/build/index.html")
	for _, file := range []string{"index.html", "service-worker.js"} {
		r.StaticFile(file, "frontend/build/"+file)
	}
}
