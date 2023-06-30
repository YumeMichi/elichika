package handler

import (
	"elichika/config"
	"elichika/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FetchProfile(ctx *gin.Context) {
	signBody := utils.ReadAllText("assets/fetchProfile.json")
	resp := SignResp(ctx.GetString("ep"), signBody, config.SessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}
