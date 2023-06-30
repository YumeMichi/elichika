package handler

import (
	"elichika/config"
	"elichika/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FetchBootstrap(ctx *gin.Context) {
	signBody := utils.ReadAllText("assets/fetchBootstrap.json")
	resp := SignResp(ctx.GetString("ep"), signBody, config.SessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func GetClearedPlatformAchievement(ctx *gin.Context) {
	signBody := utils.ReadAllText("assets/getClearedPlatformAchievement.json")
	resp := SignResp(ctx.GetString("ep"), signBody, config.SessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}
