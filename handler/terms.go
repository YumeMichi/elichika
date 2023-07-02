package handler

import (
	"elichika/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Agreement(ctx *gin.Context) {
	signBody := GetUserData("agreement.json")
	resp := SignResp(ctx.GetString("ep"), signBody, config.SessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}
