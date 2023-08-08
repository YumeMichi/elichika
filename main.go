package main

import (
	"elichika/router"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	router.Router(r)

	r.Run(":"+config.Conf.Settings.Port) // listen and serve on 0.0.0.0:80 as default, or on a custom port defined in config.json
}
