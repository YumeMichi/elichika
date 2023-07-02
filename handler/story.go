package handler

import (
	"elichika/config"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/sjson"
)

func FinishStory(ctx *gin.Context) {
	signBody, _ := sjson.Set(GetUserData("finishStory.json"),
		"user_model.user_status", GetUserStatus())
	resp := SignResp(ctx.GetString("ep"), signBody, config.SessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func FinishStoryMain(ctx *gin.Context) {
	signBody, _ := sjson.Set(GetUserData("finishUserStoryMain.json"),
		"user_model_diff.user_status", GetUserStatus())
	resp := SignResp(ctx.GetString("ep"), signBody, config.SessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func FinishStoryLinkage(ctx *gin.Context) {
	signBody, _ := sjson.Set(GetUserData("finishStoryLinkage.json"),
		"user_model_diff.user_status", GetUserStatus())
	resp := SignResp(ctx.GetString("ep"), signBody, config.SessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}
