package handler

import (
	"elichika/config"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

func FetchEmblem(ctx *gin.Context) {
	signBody, _ := sjson.Set(GetUserData("fetchEmblem.json"),
		"user_model.user_status", GetUserStatus())
	resp := SignResp(ctx.GetString("ep"), signBody, config.SessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func ActivateEmblem(ctx *gin.Context) {
	reqBody := ctx.GetString("reqBody")

	var emblemId int64
	gjson.Parse(reqBody).ForEach(func(key, value gjson.Result) bool {
		if value.Get("emblem_master_id").String() != "" {
			emblemId = value.Get("emblem_master_id").Int()

			SetUserData("userStatus.json", "emblem_id", emblemId)

			return false
		}
		return true
	})

	signBody := GetUserData("activateEmblem.json")
	signBody, _ = sjson.Set(signBody, "user_model.user_status", GetUserStatus())
	signBody, _ = sjson.Set(signBody, "user_model.user_status", GetUserStatus())
	signBody, _ = sjson.Set(signBody, "user_model.user_status.emblem_id", emblemId)
	resp := SignResp(ctx.GetString("ep"), signBody, config.SessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}
