package handler

import (
	"elichika/config"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/sjson"
)

func SaveUserNaviVoice(ctx *gin.Context) {
	signBody, _ := sjson.Set(GetData("saveUserNaviVoice.json"),
		"user_model.user_status", GetUserStatus())
	resp := SignResp(ctx.GetString("ep"), signBody, config.SessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}
