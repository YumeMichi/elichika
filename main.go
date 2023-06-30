package main

import (
	"elichika/router"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	router.Router(r)

	r.Run(":80") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
