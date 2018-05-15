package main

import (
	cfg "github.com/dxcenter/chess/config"
	db "github.com/dxcenter/chess/db"
	"github.com/dxcenter/chess/models"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg.Reload()

	myDb := db.GetDB(cfg.Get().MyDb)
	models.Init(myDb)

	r := gin.Default()
	setupRouter(r)
	r.Run()
}
