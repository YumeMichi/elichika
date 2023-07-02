package handler

import (
	"elichika/config"
	"elichika/database"
	"elichika/model"
	"elichika/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

func SaveDeckAll(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0]
	// fmt.Println(reqBody.String())

	req := model.SaveDeckReq{}
	decoder := json.NewDecoder(strings.NewReader(reqBody.String()))
	decoder.UseNumber()
	err := decoder.Decode(&req)
	CheckErr(err)
	// fmt.Println("Raw:", req.SquadDict)

	liveDeckInfo := GetUserData("liveDeck.json")
	keyDeckName := fmt.Sprintf("user_live_deck_by_id.%d.name.dot_under_text", req.DeckID*2-1)
	// fmt.Println(keyDeckName)
	deckName := gjson.Parse(liveDeckInfo).Get(keyDeckName).String()
	// fmt.Println("deckName:", deckName)

	if req.CardWithSuit[1] == 0 {
		req.CardWithSuit[1] = req.CardWithSuit[0]
	}
	if req.CardWithSuit[3] == 0 {
		req.CardWithSuit[3] = req.CardWithSuit[2]
	}
	if req.CardWithSuit[5] == 0 {
		req.CardWithSuit[5] = req.CardWithSuit[4]
	}
	if req.CardWithSuit[7] == 0 {
		req.CardWithSuit[7] = req.CardWithSuit[6]
	}
	if req.CardWithSuit[9] == 0 {
		req.CardWithSuit[9] = req.CardWithSuit[8]
	}
	if req.CardWithSuit[11] == 0 {
		req.CardWithSuit[11] = req.CardWithSuit[10]
	}
	if req.CardWithSuit[13] == 0 {
		req.CardWithSuit[13] = req.CardWithSuit[12]
	}
	if req.CardWithSuit[15] == 0 {
		req.CardWithSuit[15] = req.CardWithSuit[14]
	}
	if req.CardWithSuit[17] == 0 {
		req.CardWithSuit[17] = req.CardWithSuit[16]
	}

	deckInfo := model.DeckInfo{
		UserLiveDeckID: req.DeckID,
		Name: model.DeckName{
			DotUnderText: deckName,
		},
		CardMasterID1: req.CardWithSuit[0],
		CardMasterID2: req.CardWithSuit[2],
		CardMasterID3: req.CardWithSuit[4],
		CardMasterID4: req.CardWithSuit[6],
		CardMasterID5: req.CardWithSuit[8],
		CardMasterID6: req.CardWithSuit[10],
		CardMasterID7: req.CardWithSuit[12],
		CardMasterID8: req.CardWithSuit[14],
		CardMasterID9: req.CardWithSuit[16],
		SuitMasterID1: req.CardWithSuit[1],
		SuitMasterID2: req.CardWithSuit[3],
		SuitMasterID3: req.CardWithSuit[5],
		SuitMasterID4: req.CardWithSuit[7],
		SuitMasterID5: req.CardWithSuit[9],
		SuitMasterID6: req.CardWithSuit[11],
		SuitMasterID7: req.CardWithSuit[13],
		SuitMasterID8: req.CardWithSuit[15],
		SuitMasterID9: req.CardWithSuit[17],
	}
	// fmt.Println(deckInfo)

	keyLiveDeck := fmt.Sprintf("user_live_deck_by_id.%d", req.DeckID*2-1)
	SetUserData("liveDeck.json", keyLiveDeck, deckInfo)

	deckInfoRes := []model.AsResp{}
	deckInfoRes = append(deckInfoRes, req.DeckID)
	deckInfoRes = append(deckInfoRes, deckInfo)

	partyInfoRes := []model.AsResp{}
	for k, v := range req.SquadDict {
		if k%2 == 0 {
			partyId, err := v.(json.Number).Int64()
			if err != nil {
				panic(err)
			}
			// fmt.Println("Party ID:", partyId)

			rDictInfo, err := json.Marshal(req.SquadDict[k+1])
			CheckErr(err)

			dictInfo := model.DeckSquadDict{}
			if err = json.Unmarshal(rDictInfo, &dictInfo); err != nil {
				panic(err)
			}
			// fmt.Println("Party Info:", dictInfo)

			roleIds := []int{}
			err = MainEng.Table("m_card").
				Where("id IN (?,?,?)", dictInfo.CardMasterIds[0], dictInfo.CardMasterIds[1], dictInfo.CardMasterIds[2]).
				Cols("role").Find(&roleIds)
			CheckErr(err)
			// fmt.Println("roleIds:", roleIds)

			var partyIcon int
			var partyName string
			// 脑残逻辑部分
			exists, err := MainEng.Table("m_live_party_name").
				Where("role_1 = ? AND role_2 = ? AND role_3 = ?", roleIds[0], roleIds[1], roleIds[2]).
				Cols("name,live_party_icon").Get(&partyName, &partyIcon)
			CheckErr(err)
			if !exists {
				exists, err = MainEng.Table("m_live_party_name").
					Where("role_1 = ? AND role_2 = ? AND role_3 = ?", roleIds[0], roleIds[2], roleIds[1]).
					Cols("name,live_party_icon").Get(&partyName, &partyIcon)
				CheckErr(err)
				if !exists {
					exists, err = MainEng.Table("m_live_party_name").
						Where("role_1 = ? AND role_2 = ? AND role_3 = ?", roleIds[1], roleIds[0], roleIds[2]).
						Cols("name,live_party_icon").Get(&partyName, &partyIcon)
					CheckErr(err)
					if !exists {
						exists, err = MainEng.Table("m_live_party_name").
							Where("role_1 = ? AND role_2 = ? AND role_3 = ?", roleIds[1], roleIds[2], roleIds[0]).
							Cols("name,live_party_icon").Get(&partyName, &partyIcon)
						CheckErr(err)
						if !exists {
							exists, err = MainEng.Table("m_live_party_name").
								Where("role_1 = ? AND role_2 = ? AND role_3 = ?", roleIds[2], roleIds[0], roleIds[1]).
								Cols("name,live_party_icon").Get(&partyName, &partyIcon)
							CheckErr(err)
							if !exists {
								exists, err = MainEng.Table("m_live_party_name").
									Where("role_1 = ? AND role_2 = ? AND role_3 = ?", roleIds[2], roleIds[1], roleIds[0]).
									Cols("name,live_party_icon").Get(&partyName, &partyIcon)
								CheckErr(err)
								if !exists {
									panic("Fuck you!")
								}
							}
						}
					}
				}
			}

			var realPartyName string
			_, err = MainEng.Table("m_dictionary").Where("id = ?", strings.ReplaceAll(partyName, "k.", "")).Cols("message").Get(&realPartyName)
			CheckErr(err)

			partyInfo := model.PartyInfo{
				PartyID:        int(partyId),
				UserLiveDeckID: req.DeckID,
				Name: model.PartyName{
					DotUnderText: realPartyName,
				},
				IconMasterID:     partyIcon,
				CardMasterID1:    dictInfo.CardMasterIds[0],
				CardMasterID2:    dictInfo.CardMasterIds[1],
				CardMasterID3:    dictInfo.CardMasterIds[2],
				UserAccessoryID1: dictInfo.UserAccessoryIds[0],
				UserAccessoryID2: dictInfo.UserAccessoryIds[1],
				UserAccessoryID3: dictInfo.UserAccessoryIds[2],
			}
			// fmt.Println(partyInfo)

			keyLiveParty := fmt.Sprintf("user_live_party_by_id.%d", (req.DeckID-1)*6+(k+1))
			SetUserData("liveDeck.json", keyLiveParty, partyInfo)
			// fmt.Println(keyLiveParty, partyName)

			partyInfoRes = append(partyInfoRes, partyId)
			partyInfoRes = append(partyInfoRes, partyInfo)
		}
	}

	respBody := GetUserData("saveDeckAll.json")
	respBody, _ = sjson.Set(respBody, "user_model.user_status", GetUserStatus())
	respBody, _ = sjson.Set(respBody, "user_model.user_live_deck_by_id", deckInfoRes)
	respBody, _ = sjson.Set(respBody, "user_model.user_live_party_by_id", partyInfoRes)
	resp := SignResp(ctx.GetString("ep"), respBody, config.SessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func FetchLiveMusicSelect(ctx *gin.Context) {
	now := time.Now()
	year, month, day := now.Year(), now.Month(), now.Day()
	tomorrow := time.Date(year, month, day+1, 0, 0, 0, 0, now.Location()).Unix()
	weekday := int(now.Weekday())

	liveDailyList := []model.LiveDaily{}
	err := MainEng.Table("m_live_daily").Where("weekday = ?", weekday).Cols("id,live_id").Find(&liveDailyList)
	CheckErr(err)
	for k := range liveDailyList {
		liveDailyList[k].EndAt = int(tomorrow)
		liveDailyList[k].RemainingPlayCount = 5
		liveDailyList[k].RemainingRecoveryCount = 9
	}

	signBody := GetUserData("fetchLiveMusicSelect.json")
	signBody, _ = sjson.Set(signBody, "weekday_state.weekday", weekday)
	signBody, _ = sjson.Set(signBody, "weekday_state.next_weekday_at", tomorrow)
	signBody, _ = sjson.Set(signBody, "live_daily_list", liveDailyList)
	signBody, _ = sjson.Set(signBody, "user_model_diff.user_status", GetUserStatus())
	resp := SignResp(ctx.GetString("ep"), signBody, config.SessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func FetchLivePartners(ctx *gin.Context) {
	signBody := GetUserData("fetchLivePartners.json")
	resp := SignResp(ctx.GetString("ep"), signBody, config.SessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func FetchLiveDeckSelect(ctx *gin.Context) {
	signBody := GetUserData("fetchLiveDeckSelect.json")
	resp := SignResp(ctx.GetString("ep"), signBody, config.SessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func LiveStart(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0]
	// fmt.Println(reqBody.String())

	liveStartReq := model.LiveStartReq{}
	if err := json.Unmarshal([]byte(reqBody.String()), &liveStartReq); err != nil {
		panic(err)
	}
	// fmt.Println(liveStartReq)

	var cardInfo string
	partnerResp := gjson.Parse(GetUserData("fetchLivePartners.json")).Get("partner_select_state.live_partners")
	partnerResp.ForEach(func(k, v gjson.Result) bool {
		userId := v.Get("user_id").Int()
		if userId == int64(liveStartReq.PartnerUserID) {
			v.Get("card_by_category").ForEach(func(kk, vv gjson.Result) bool {
				if vv.IsObject() {
					cardId := vv.Get("card_master_id").Int()
					if cardId == int64(liveStartReq.PartnerCardMasterID) {
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

	// 保存请求包因为 /live/finish 接口的响应包里有部分字段不在该接口的请求包里
	liveId := time.Now().UnixNano()
	liveIdStr := strconv.Itoa(int(liveId))
	err := database.LevelDb.Put([]byte("live_"+liveIdStr), []byte(reqBody.String()))
	CheckErr(err)

	liveDifficultyId := strconv.Itoa(liveStartReq.LiveDifficultyID)
	liveNotes := utils.ReadAllText("assets/stages/" + liveDifficultyId + ".json")
	if liveNotes == "" {
		panic("歌曲情报信息不存在！")
	}

	var liveNotesRes map[string]any
	if err = json.Unmarshal([]byte(liveNotes), &liveNotesRes); err != nil {
		panic(err)
	}

	var partnerInfo any
	if cardInfo != "" {
		var info map[string]any
		if err = json.Unmarshal([]byte(cardInfo), &info); err != nil {
			panic(err)
		}
		partnerInfo = info
	} else {
		partnerInfo = nil
	}

	liveStartResp := GetUserData("liveStart.json")
	liveStartResp, _ = sjson.Set(liveStartResp, "live.live_id", liveId)
	liveStartResp, _ = sjson.Set(liveStartResp, "live.deck_id", liveStartReq.DeckID)
	liveStartResp, _ = sjson.Set(liveStartResp, "live.live_stage", liveNotesRes)
	liveStartResp, _ = sjson.Set(liveStartResp, "live.live_partner_card", partnerInfo)
	liveStartResp, _ = sjson.Set(liveStartResp, "user_model_diff.user_status", GetUserStatus())
	liveStartResp, _ = sjson.Set(liveStartResp, "user_model_diff.user_status.latest_live_deck_id", liveStartReq.DeckID)
	liveStartResp, _ = sjson.Set(liveStartResp, "user_model_diff.user_status.last_live_difficulty_id", liveStartReq.LiveDifficultyID)
	resp := SignResp(ctx.GetString("ep"), liveStartResp, config.SessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func LiveFinish(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0]
	// fmt.Println(reqBody.String())

	var cardMasterId, maxVolt, skillCount, appealCount int64
	liveFinishReq := gjson.Parse(reqBody.String())
	liveFinishReq.Get("live_score.card_stat_dict").ForEach(func(key, value gjson.Result) bool {
		if value.IsObject() {
			volt := value.Get("got_voltage").Int()
			if volt > maxVolt {
				maxVolt = volt

				cardMasterId = value.Get("card_master_id").Int()
				skillCount = value.Get("skill_triggered_count").Int()
				appealCount = value.Get("appeal_count").Int()
			}
		}
		return true
	})

	mvpInfo := model.MvpInfo{
		CardMasterID:        cardMasterId,
		GetVoltage:          maxVolt,
		SkillTriggeredCount: skillCount,
		AppealCount:         appealCount,
	}

	liveId := liveFinishReq.Get("live_id").String()
	res, err := database.LevelDb.Get([]byte("live_" + liveId))
	CheckErr(err)

	liveStartReq := model.LiveStartReq{}
	if err := json.Unmarshal(res, &liveStartReq); err != nil {
		panic(err)
	}
	// fmt.Println("liveStartReq:", liveStartReq)

	var partnerInfo any
	if liveStartReq.PartnerUserID != 0 {
		info := model.LivePartnerInfo{
			LastPlayedAt:                        time.Now().Unix(),
			RecommendCardMasterID:               liveStartReq.PartnerCardMasterID,
			RecommendCardLevel:                  1,
			IsRecommendCardImageAwaken:          true,
			IsRecommendCardAllTrainingActivated: true,
			IsNew:                               false,
			FriendApprovedAt:                    nil,
			RequestStatus:                       3,
			IsRequestPending:                    false,
		}
		partnerResp := gjson.Parse(GetUserData("fetchLivePartners.json")).Get("partner_select_state.live_partners")
		partnerResp.ForEach(func(k, v gjson.Result) bool {
			userId := v.Get("user_id").Int()
			if userId == int64(liveStartReq.PartnerUserID) {
				info.UserID = int(userId)
				info.Name.DotUnderText = v.Get("name.dot_under_text").String()
				info.Rank = int(v.Get("rank").Int())
				info.EmblemID = int(v.Get("emblem_id").Int())
				info.IntroductionMessage.DotUnderText = v.Get("introduction_message.dot_under_text").String()
			}
			return true
		})
		partnerInfo = info
	} else {
		partnerInfo = nil
	}

	liveResult := model.LiveResultAchievementStatus{
		ClearCount:       1,
		GotVoltage:       liveFinishReq.Get("live_score.current_score").Int(),
		RemainingStamina: liveFinishReq.Get("live_score.remaining_stamina").Int(),
	}

	liveFinishResp := GetUserData("liveFinish.json")
	liveFinishResp, _ = sjson.Set(liveFinishResp, "live_result.live_difficulty_master_id", liveStartReq.LiveDifficultyID)
	liveFinishResp, _ = sjson.Set(liveFinishResp, "live_result.live_deck_id", liveStartReq.DeckID)
	liveFinishResp, _ = sjson.Set(liveFinishResp, "live_result.mvp", mvpInfo)
	liveFinishResp, _ = sjson.Set(liveFinishResp, "live_result.partner", partnerInfo)
	liveFinishResp, _ = sjson.Set(liveFinishResp, "live_result.live_result_achievement_status", liveResult)
	liveFinishResp, _ = sjson.Set(liveFinishResp, "live_result.voltage", liveFinishReq.Get("live_score.current_score").Int())
	liveFinishResp, _ = sjson.Set(liveFinishResp, "live_result.last_best_voltage", liveFinishReq.Get("live_score.current_score").Int())
	liveFinishResp, _ = sjson.Set(liveFinishResp, "live_result.before_user_exp", GetUserStatus()["exp"].(float64))
	liveFinishResp, _ = sjson.Set(liveFinishResp, "live_result.gain_user_exp", 0)
	liveFinishResp, _ = sjson.Set(liveFinishResp, "user_model_diff.user_status", GetUserStatus())
	liveFinishResp, _ = sjson.Set(liveFinishResp, "user_model_diff.user_status.latest_live_deck_id", liveStartReq.DeckID)
	liveFinishResp, _ = sjson.Set(liveFinishResp, "user_model_diff.user_status.last_live_difficulty_id", liveStartReq.LiveDifficultyID)
	resp := SignResp(ctx.GetString("ep"), liveFinishResp, config.SessionKey)
	// fmt.Println(resp)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func LiveMvStart(ctx *gin.Context) {
	signBody, _ := sjson.Set(GetUserData("liveMvStart.json"),
		"user_model_diff.user_status", GetUserStatus())
	resp := SignResp(ctx.GetString("ep"), signBody, config.SessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func LiveMvSaveDeck(ctx *gin.Context) {
	reqBody := ctx.GetString("reqBody")
	// fmt.Println(reqBody)

	reqData := gjson.Parse(reqBody).Array()[0]
	// fmt.Println(reqData)

	saveReq := model.LiveSaveDeckReq{}
	err := json.Unmarshal([]byte(reqData.String()), &saveReq)
	if err != nil {
		panic(err)
	}
	// fmt.Println(saveReq)

	userLiveMvDeckInfo := model.UserLiveMvDeckInfo{
		LiveMasterID: saveReq.LiveMasterID,
	}

	memberIds := map[int]int{}
	for k, v := range saveReq.MemberMasterIDByPos {
		if k%2 == 0 {
			memberId := saveReq.MemberMasterIDByPos[k+1]
			memberIds[v] = memberId

			switch v {
			case 1:
				userLiveMvDeckInfo.MemberMasterID1 = memberId
			case 2:
				userLiveMvDeckInfo.MemberMasterID2 = memberId
			case 3:
				userLiveMvDeckInfo.MemberMasterID3 = memberId
			case 4:
				userLiveMvDeckInfo.MemberMasterID4 = memberId
			case 5:
				userLiveMvDeckInfo.MemberMasterID5 = memberId
			case 6:
				userLiveMvDeckInfo.MemberMasterID6 = memberId
			case 7:
				userLiveMvDeckInfo.MemberMasterID7 = memberId
			case 8:
				userLiveMvDeckInfo.MemberMasterID8 = memberId
			case 9:
				userLiveMvDeckInfo.MemberMasterID9 = memberId
			case 10:
				userLiveMvDeckInfo.MemberMasterID10 = memberId
			case 11:
				userLiveMvDeckInfo.MemberMasterID11 = memberId
			case 12:
				userLiveMvDeckInfo.MemberMasterID12 = memberId
			}
		}
	}
	// fmt.Println(memberIds)

	suitIds := map[int]int{}
	for k, v := range saveReq.SuitMasterIDByPos {
		if k%2 == 0 {
			suitId := saveReq.SuitMasterIDByPos[k+1]
			suitIds[v] = suitId

			switch v {
			case 1:
				userLiveMvDeckInfo.SuitMasterID1 = suitId
			case 2:
				userLiveMvDeckInfo.SuitMasterID2 = suitId
			case 3:
				userLiveMvDeckInfo.SuitMasterID3 = suitId
			case 4:
				userLiveMvDeckInfo.SuitMasterID4 = suitId
			case 5:
				userLiveMvDeckInfo.SuitMasterID5 = suitId
			case 6:
				userLiveMvDeckInfo.SuitMasterID6 = suitId
			case 7:
				userLiveMvDeckInfo.SuitMasterID7 = suitId
			case 8:
				userLiveMvDeckInfo.SuitMasterID8 = suitId
			case 9:
				userLiveMvDeckInfo.SuitMasterID9 = suitId
			case 10:
				userLiveMvDeckInfo.SuitMasterID10 = suitId
			case 11:
				userLiveMvDeckInfo.SuitMasterID11 = suitId
			case 12:
				userLiveMvDeckInfo.SuitMasterID12 = suitId
			}
		}
	}
	// fmt.Println(suitIds)

	var userLiveMvDeckCustomByID []any
	userLiveMvDeckCustomByID = append(userLiveMvDeckCustomByID, saveReq.LiveMasterID)
	userLiveMvDeckCustomByID = append(userLiveMvDeckCustomByID, userLiveMvDeckInfo)
	// fmt.Println(userLiveMvDeckCustomByID)

	viewStatusIds := map[int]int{}
	for k, v := range saveReq.ViewStatusByPos {
		if k%2 == 0 {
			viewStatusId := saveReq.ViewStatusByPos[k+1]
			viewStatusIds[v] = viewStatusId
		}
	}
	// fmt.Println(viewStatusIds)

	var userMemberByMemberID []any
	for k, v := range memberIds {
		userMemberByMemberID = append(userMemberByMemberID, v)
		userMemberByMemberID = append(userMemberByMemberID, model.UserMemberInfo{
			MemberMasterID:           v,
			CustomBackgroundMasterID: 103506600,
			SuitMasterID:             suitIds[k],
			LovePoint:                0,
			LovePointLimit:           999999,
			LoveLevel:                1,
			ViewStatus:               viewStatusIds[k],
			IsNew:                    true,
		})
	}
	// fmt.Println(userMemberByMemberID)

	signBody := GetUserData("liveMvSaveDeck.json")
	signBody, _ = sjson.Set(signBody, "user_model.user_status", GetUserStatus())
	signBody, _ = sjson.Set(signBody, "user_model.user_member_by_member_id", userMemberByMemberID)
	signBody, _ = sjson.Set(signBody, "user_model.user_live_mv_deck_custom_by_id", userLiveMvDeckCustomByID)

	resp := SignResp(ctx.GetString("ep"), string(signBody), config.SessionKey)
	// fmt.Println(resp)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}
