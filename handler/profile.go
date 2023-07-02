package handler

import (
	"elichika/config"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

func FetchProfile(ctx *gin.Context) {
	userInfo := gjson.Parse(GetUserData("userStatus.json"))
	signBody := GetUserData("fetchProfile.json")
	signBody, _ = sjson.Set(signBody, "profile_info.basic_info.name.dot_under_text",
		userInfo.Get("name.dot_under_text").String())
	signBody, _ = sjson.Set(signBody, "profile_info.basic_info.introduction_message.dot_under_text",
		userInfo.Get("message.dot_under_text").String())
	signBody, _ = sjson.Set(signBody, "profile_info.basic_info.emblem_id",
		userInfo.Get("emblem_id").Int())
	resp := SignResp(ctx.GetString("ep"), signBody, config.SessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func SetProfile(ctx *gin.Context) {
	reqBody := ctx.GetString("reqBody")
	// fmt.Println(reqBody)

	req := gjson.Parse(reqBody).Array()[0]
	if req.Get("name").String() != "" {
		SetUserData("userStatus.json", "name.dot_under_text",
			gjson.Parse(reqBody).Array()[0].Get("name").String())
	} else if req.Get("nickname").String() != "" {
		SetUserData("userStatus.json", "nickname.dot_under_text",
			gjson.Parse(reqBody).Array()[0].Get("nickname").String())
	} else if req.Get("message").String() != "" {
		SetUserData("userStatus.json", "message.dot_under_text",
			gjson.Parse(reqBody).Array()[0].Get("message").String())
	}

	signBody, _ := sjson.Set(GetUserData("setProfile.json"),
		"user_model.user_status", GetUserStatus())
	resp := SignResp(ctx.GetString("ep"), signBody, config.SessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}
