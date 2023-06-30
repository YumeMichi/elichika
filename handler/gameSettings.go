package handler

import (
	"elichika/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdatePushNotificationSettings(ctx *gin.Context) {
	resp := SignResp(ctx.GetString("ep"), "{}", config.SessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}
