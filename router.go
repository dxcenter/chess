package main

import (
	jwt "github.com/dxcenter/chess/jwt"
	m "github.com/dxcenter/chess/serverMethods"
	mw "github.com/dxcenter/chess/serverMethods/middleware"
	"github.com/gin-gonic/gin"
)

func setupRouter(r *gin.Engine) {
	setupJsonRouter(r)
	setupFrontendRouter(r)
}

func setupJsonRouter(r *gin.Engine) {
	// JWT
	jwtMiddleware := jwt.GetJwtMiddleware()
	authed := r.Group("/")
	authed.Use(jwtMiddleware.MiddlewareFunc()) // require to be authed
	authed.Use(mw.Authed)                      // some routines for an already authed
	r.POST("/auth.json", jwtMiddleware.LoginHandler)
	r.POST("/refresh_token.json", jwtMiddleware.RefreshHandler)

	// Mix of my methods and JWT
	signUp := r.Group("/")
	signUp.Use(mw.SignUp)
	signUp.POST("/sign_up.json", jwtMiddleware.LoginHandler)

	// My methods
	r.GET("/ping.json", m.Ping)
	authed.GET("/whoami.json", m.Whoami)
	authed.GET("/players.json", m.Players)
	authed.GET("/players_pairs.json", m.PlayersPairs)
	authed.GET("/games.json", m.Games)
	authed.GET("/games/:game_id/status.json", m.GameStatus)
	authed.POST("/games.json", m.NewGame)
	//authed.DELETE("/games.json", m.EndGame)
	authed.POST("/games/:game_id/move.json", m.Move)
}

func setupFrontendRouter(r *gin.Engine) {
	r.Static("/frontend", "frontend/build")
	r.Static("/static", "frontend/build/static")
	r.Static("/css", "frontend/build/css")
	r.StaticFile("/", "frontend/build/index.html")
	//r.StaticFile("/games/:game_id", "frontend/build/index.html") // idk how to make that, yet. Use nginx :)
	r.StaticFile("/games_new", "frontend/build/index.html") // should be "/games/new" (not "/games_new"). See https://github.com/gin-gonic/gin/issues/360
	r.StaticFile("/login", "frontend/build/index.html")
	r.StaticFile("/signup", "frontend/build/index.html")
	for _, file := range []string{"index.html", "service-worker.js"} {
		r.StaticFile(file, "frontend/build/"+file)
	}
}
