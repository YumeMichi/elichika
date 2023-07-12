package middleware

import (
	"elichika/config"
	"io"
	"strings"

	"github.com/gin-gonic/gin"
)

func Common(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		panic(err)
	}
	defer ctx.Request.Body.Close()
	ctx.Set("reqBody", string(body))

	ep := strings.ReplaceAll(ctx.Request.URL.String(), "/"+config.Conf.Settings.EndPoint, "")
	ctx.Set("ep", ep)

	ctx.Next()
}
