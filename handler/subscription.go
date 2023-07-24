package handler

import (
	"elichika/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FetchSubscriptionPass(ctx *gin.Context) {
	signBody := GetData("fetchSubscriptionPass.json")
	resp := SignResp(ctx.GetString("ep"), signBody, config.SessionKey)
	// fmt.Println(resp)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}
