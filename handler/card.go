package handler

import (
	"elichika/config"
	"elichika/model"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

func UpdateCardNewFlag(ctx *gin.Context) {
	signBody, _ := sjson.Set(GetUserData("updateCardNewFlag.json"),
		"user_model_diff.user_status", GetUserStatus())
	resp := SignResp(ctx.GetString("ep"), signBody, config.SessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func ChangeIsAwakeningImage(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0]
	// fmt.Println(reqBody.String())

	req := model.CardAwakeningReq{}
	if err := json.Unmarshal([]byte(reqBody.String()), &req); err != nil {
		panic(err)
	}

	loginData := GetUserData("login.json")
	cardInfo := model.CardInfo{}
	gjson.Parse(loginData).Get("user_model.user_card_by_card_id").
		ForEach(func(key, value gjson.Result) bool {
			if value.IsObject() {
				if err := json.Unmarshal([]byte(value.String()), &cardInfo); err != nil {
					panic(err)
				}

				if cardInfo.CardMasterID == req.CardMasterID {
					cardInfo.IsAwakeningImage = req.IsAwakeningImage

					k := "user_model.user_card_by_card_id." + key.String() + ".is_awakening_image"
					SetUserData("login.json", k, req.IsAwakeningImage)

					return false
				}
			}
			return true
		})

	userCardInfo := []any{}
	userCardInfo = append(userCardInfo, cardInfo.CardMasterID)
	userCardInfo = append(userCardInfo, cardInfo)

	cardResp := GetUserData("changeIsAwakeningImage.json")
	cardResp, _ = sjson.Set(cardResp, "user_model_diff.user_status", GetUserStatus())
	cardResp, _ = sjson.Set(cardResp, "user_model_diff.user_card_by_card_id", userCardInfo)
	resp := SignResp(ctx.GetString("ep"), cardResp, config.SessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}
