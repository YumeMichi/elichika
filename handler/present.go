package handler

import (
	"elichika/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FetchPresent(ctx *gin.Context) {
	signBody := GetData("fetchPresent.json")
	resp := SignResp(ctx.GetString("ep"), signBody, config.SessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}
