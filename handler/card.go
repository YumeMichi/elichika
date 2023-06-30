package handler

import (
	"elichika/config"
	"elichika/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateCardNewFlag(ctx *gin.Context) {
	signBody := utils.ReadAllText("assets/updateCardNewFlag.json")
	resp := SignResp(ctx.GetString("ep"), signBody, config.SessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}
