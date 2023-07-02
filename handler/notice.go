package handler

import (
	"elichika/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FetchNotice(ctx *gin.Context) {
	signBody := GetUserData("fetchNotice.json")
	resp := SignResp(ctx.GetString("ep"), signBody, config.SessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}
