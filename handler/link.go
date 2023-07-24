package handler

import (
	"elichika/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Sif2DataLink(ctx *gin.Context) {
	signBody := GetData("sif2DataLink.json")
	resp := SignResp(ctx.GetString("ep"), signBody, config.SessionKey)
	// fmt.Println(resp)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}
