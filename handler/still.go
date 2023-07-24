package handler

import (
	"elichika/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FetchStill(ctx *gin.Context) {
	signBody := GetData("fetchStill.json")
	resp := SignResp(ctx.GetString("ep"), signBody, config.SessionKey)
	// fmt.Println(resp)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}
