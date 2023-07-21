package handler

import (
	"elichika/config"
	"elichika/model"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

func FetchProfile(ctx *gin.Context) {
	userId, _ := strconv.Atoi(ctx.Query("u"))
	userInfo := gjson.Parse(GetUserData("userStatus.json"))
	signBody := GetUserData("fetchProfile.json")
	signBody, _ = sjson.Set(signBody, "profile_info.basic_info.name.dot_under_text",
		userInfo.Get("name.dot_under_text").String())
	signBody, _ = sjson.Set(signBody, "profile_info.basic_info.introduction_message.dot_under_text",
		userInfo.Get("message.dot_under_text").String())
	signBody, _ = sjson.Set(signBody, "profile_info.basic_info.emblem_id",
		userInfo.Get("emblem_id").Int())
	signBody, _ = sjson.Set(signBody, "profile_info.basic_info.user_id", userId)
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

func SetRecommendCard(ctx *gin.Context) {
	reqBody := ctx.GetString("reqBody")
	// fmt.Println(reqBody)

	cardMasterId := gjson.Parse(reqBody).Array()[0].Get("card_master_id").Int()
	var cardInfo model.CardInfo
	gjson.Parse(GetUserData("userCard.json")).Get("user_card_by_card_id").ForEach(func(key, value gjson.Result) bool {
		if value.IsObject() {
			if value.Get("card_master_id").Int() == cardMasterId {
				if err := json.Unmarshal([]byte(value.String()), &cardInfo); err != nil {
					panic(err)
				}
				return false
			}
		}
		return true
	})

	SetUserData("userStatus.json", "recommend_card_master_id", cardMasterId)
	SetUserData("fetchProfile.json", "profile_info.basic_info.recommend_card_master_id", cardMasterId)
	SetUserData("fetchProfile.json", "profile_info.basic_info.is_recommend_card_image_awaken", cardInfo.IsAwakeningImage)

	signBody, _ := sjson.Set(GetUserData("setRecommendCard.json"),
		"user_model.user_status", GetUserStatus())
	resp := SignResp(ctx.GetString("ep"), signBody, config.SessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}
