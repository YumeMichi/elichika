package main

import (
	"elichika/router"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	router.Router(r)

	r.Run(":8080") // listen and serve on 127.0.0.1:8080 (for windows "localhost:8080")
}
