package middleware

import (
	"elichika/handler"
	"io"

	"github.com/gin-gonic/gin"
)

func Common(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		panic(err)
	}
	defer ctx.Request.Body.Close()
	ctx.Set("reqBody", string(body))

	lang, _ := ctx.GetQuery("l")
	if lang == "" {
		handler.IsGlobal = false
		handler.MasterVersion = "b66ec2295e9a00aa"
		handler.StartUpKey = "5f7IZY1QrAX0D49g"
	} else {
		handler.IsGlobal = true
		handler.MasterVersion = "2d61e7b4e89961c7"
		handler.StartUpKey = "TxQFwgNcKDlesb93"
	}

	ctx.Set("ep", ctx.Request.URL.String())

	ctx.Next()
}
