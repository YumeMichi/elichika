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

	loginData := GetUserData("userCard.json")
	cardInfo := model.CardInfo{}
	gjson.Parse(loginData).Get("user_card_by_card_id").
		ForEach(func(key, value gjson.Result) bool {
			if value.IsObject() {
				if err := json.Unmarshal([]byte(value.String()), &cardInfo); err != nil {
					panic(err)
				}

				if cardInfo.CardMasterID == req.CardMasterID {
					cardInfo.IsAwakeningImage = req.IsAwakeningImage

					k := "user_card_by_card_id." + key.String() + ".is_awakening_image"
					SetUserData("userCard.json", k, req.IsAwakeningImage)

					return false
				}
			}
			return true
		})

	userCardInfo := []any{}
	userCardInfo = append(userCardInfo, cardInfo.CardMasterID)
	userCardInfo = append(userCardInfo, cardInfo)

	// Update user profile
	cardMasterId := gjson.Parse(GetUserData("fetchProfile.json")).Get("profile_info.basic_info.recommend_card_master_id").Int()
	if cardMasterId == int64(req.CardMasterID) {
		SetUserData("fetchProfile.json", "profile_info.basic_info.is_recommend_card_image_awaken", req.IsAwakeningImage)
	}

	cardResp := GetUserData("changeIsAwakeningImage.json")
	cardResp, _ = sjson.Set(cardResp, "user_model_diff.user_status", GetUserStatus())
	cardResp, _ = sjson.Set(cardResp, "user_model_diff.user_card_by_card_id", userCardInfo)
	resp := SignResp(ctx.GetString("ep"), cardResp, config.SessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func ChangeFavorite(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0]
	// fmt.Println(reqBody.String())

	req := model.CardFavoriteReq{}
	if err := json.Unmarshal([]byte(reqBody.String()), &req); err != nil {
		panic(err)
	}

	cardData := GetUserData("userCard.json")
	cardInfo := model.CardInfo{}
	gjson.Parse(cardData).Get("user_card_by_card_id").
		ForEach(func(key, value gjson.Result) bool {
			if value.IsObject() {
				if err := json.Unmarshal([]byte(value.String()), &cardInfo); err != nil {
					panic(err)
				}

				if cardInfo.CardMasterID == req.CardMasterID {
					cardInfo.IsFavorite = req.IsFavorite

					k := "user_card_by_card_id." + key.String() + ".is_favorite"
					SetUserData("userCard.json", k, req.IsFavorite)

					return false
				}
			}
			return true
		})

	userCardInfo := []any{}
	userCardInfo = append(userCardInfo, cardInfo.CardMasterID)
	userCardInfo = append(userCardInfo, cardInfo)

	cardResp := GetUserData("changeFavorite.json")
	cardResp, _ = sjson.Set(cardResp, "user_model_diff.user_status", GetUserStatus())
	cardResp, _ = sjson.Set(cardResp, "user_model_diff.user_card_by_card_id", userCardInfo)
	resp := SignResp(ctx.GetString("ep"), cardResp, config.SessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func GetOtherUserCard(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0]
	// fmt.Println(reqBody.String())

	userCardReq := model.UserCardReq{}
	if err := json.Unmarshal([]byte(reqBody.String()), &userCardReq); err != nil {
		panic(err)
	}
	// fmt.Println(liveStartReq)

	var newUserCardInfo model.NewCardInfo
	var cardInfo string
	partnerList := gjson.Parse(GetUserData("fetchLivePartners.json")).Get("partner_select_state.live_partners")
	partnerList.ForEach(func(k, v gjson.Result) bool {
		userId := v.Get("user_id").Int()
		if userId == userCardReq.UserID {
			v.Get("card_by_category").ForEach(func(kk, vv gjson.Result) bool {
				if vv.IsObject() {
					cardId := vv.Get("card_master_id").Int()
					if cardId == userCardReq.CardMasterID {
						cardInfo = vv.String()
						// fmt.Println(cardInfo)
						return false
					}
				}
				return true
			})
			return false
		}
		return true
	})

	if err := json.Unmarshal([]byte(cardInfo), &newUserCardInfo); err != nil {
		panic(err)
	}

	userCardResp := GetUserData("getOtherUserCard.json")
	userCardResp, _ = sjson.Set(userCardResp, "other_user_card", newUserCardInfo)
	resp := SignResp(ctx.GetString("ep"), userCardResp, config.SessionKey)
	// fmt.Println(resp)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func FetchTrainingTree(ctx *gin.Context) {
	signBody := GetUserData("fetchTrainingTree.json")
	resp := SignResp(ctx.GetString("ep"), signBody, config.SessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}
