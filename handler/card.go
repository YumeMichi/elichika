package handler

import (
	"elichika/config"
	"elichika/model"
	"elichika/utils"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

func UpdateCardNewFlag(ctx *gin.Context) {
	signBody, _ := sjson.Set(utils.ReadAllText("assets/updateCardNewFlag.json"),
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

	cardInfo := model.CardInfo{}
	gjson.Parse(utils.ReadAllText("assets/login.json")).Get("user_model.user_card_by_card_id").
		ForEach(func(key, value gjson.Result) bool {
			if value.IsObject() {
				if err := json.Unmarshal([]byte(value.String()), &cardInfo); err != nil {
					panic(err)
				}

				if cardInfo.CardMasterID == req.CardMasterID {
					cardInfo.IsAwakeningImage = req.IsAwakeningImage

					return false
				}
			}
			return true
		})

	userCardInfo := []any{}
	userCardInfo = append(userCardInfo, cardInfo.CardMasterID)
	userCardInfo = append(userCardInfo, cardInfo)

	cardResp := utils.ReadAllText("assets/changeIsAwakeningImage.json")
	cardResp, _ = sjson.Set(cardResp, "user_model_diff.user_status", GetUserStatus())
	cardResp, _ = sjson.Set(cardResp, "user_model_diff.user_card_by_card_id", userCardInfo)
	resp := SignResp(ctx.GetString("ep"), cardResp, config.SessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}
