package handler

import (
	"elichika/config"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/sjson"
)

func FetchBootstrap(ctx *gin.Context) {
	signBody, _ := sjson.Set(GetUserData("fetchBootstrap.json"),
		"user_model_diff.user_status", GetUserStatus())
	resp := SignResp(ctx.GetString("ep"), signBody, config.SessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func GetClearedPlatformAchievement(ctx *gin.Context) {
	signBody := GetUserData("getClearedPlatformAchievement.json")
	resp := SignResp(ctx.GetString("ep"), signBody, config.SessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}
