package main

import (
	"elichika/config"
	"elichika/patcher"
	"elichika/router"

	"github.com/gin-gonic/gin"
)

func main() {
	patcher.ApkPatcher()

	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	router.Router(r)

	r.Run(":" + config.Conf.Settings.ListenPort)
}
