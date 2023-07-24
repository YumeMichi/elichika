package handler

import (
	"elichika/config"
	"elichika/model"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

func ExecuteLesson(ctx *gin.Context) {
	reqBody := ctx.GetString("reqBody")
	// fmt.Println(reqBody)

	req := gjson.Parse(reqBody).Array()[0]
	deckId := req.Get("selected_deck_id").Int()

	var deckInfo string
	var actionList []model.LessonMenuAction
	gjson.Parse(GetUserData("lessonDeck.json")).Get("user_lesson_deck_by_id").ForEach(func(key, value gjson.Result) bool {
		if value.IsObject() && value.Get("user_lesson_deck_id").Int() == deckId {
			deckInfo = value.String()
			// fmt.Println("Deck Info:", deckInfo)

			gjson.Parse(deckInfo).ForEach(func(kk, vv gjson.Result) bool {
				// fmt.Printf("kk: %s, vv: %s\n", kk.String(), vv.String())
				if strings.Contains(kk.String(), "card_master_id") {
					actionList = append(actionList, model.LessonMenuAction{
						CardMasterID:                  vv.Int(),
						Position:                      0,
						IsAddedPassiveSkill:           true,
						IsAddedSpecialPassiveSkill:    true,
						IsRankupedPassiveSkill:        true,
						IsRankupedSpecialPassiveSkill: true,
						IsPromotedSkill:               true,
						MaxRarity:                     4,
						UpCount:                       1,
					})
				}
				return true
			})
			return false
		}
		return true
	})
	// fmt.Println(actionList)

	SetUserData("userStatus.json", "main_lesson_deck_id", deckId)

	signBody := GetData("executeLesson.json")
	signBody, _ = sjson.Set(signBody, "lesson_menu_actions.1", actionList)
	signBody, _ = sjson.Set(signBody, "lesson_menu_actions.3", actionList)
	signBody, _ = sjson.Set(signBody, "lesson_menu_actions.5", actionList)
	signBody, _ = sjson.Set(signBody, "lesson_menu_actions.7", actionList)
	signBody, _ = sjson.Set(signBody, "user_model_diff.user_status", GetUserStatus())
	resp := SignResp(ctx.GetString("ep"), signBody, config.SessionKey)
	// fmt.Println(resp)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func ResultLesson(ctx *gin.Context) {
	userData := GetUserStatus()
	signBody, _ := sjson.Set(GetData("resultLesson.json"),
		"user_model_diff.user_status", userData)
	signBody, _ = sjson.Set(signBody, "selected_deck_id", userData["main_lesson_deck_id"])
	resp := SignResp(ctx.GetString("ep"), signBody, config.SessionKey)
	// fmt.Println(resp)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func SkillEditResult(ctx *gin.Context) {
	reqBody := ctx.GetString("reqBody")
	// fmt.Println(reqBody)

	req := gjson.Parse(reqBody).Array()[0]

	var cardList []any
	index := 1
	cardData := GetUserData("userCard.json")
	cardInfo := gjson.Parse(cardData).Get("user_card_by_card_id")
	cardInfo.ForEach(func(key, value gjson.Result) bool {
		if value.IsObject() {
			if index > 9 {
				return false
			}
			// fmt.Println("cardInfo:", value.String())

			skillList := req.Get("selected_skill_ids")
			skillList.ForEach(func(kk, vv gjson.Result) bool {
				if kk.Int()%2 == 0 && vv.Int() == value.Get("card_master_id").Int() {
					skill := skillList.Get(fmt.Sprintf("%d", kk.Int()+1))
					skill.ForEach(func(kkk, vvv gjson.Result) bool {
						skillIdKey := fmt.Sprintf("user_card_by_card_id.%s.additional_passive_skill_%d_id", key.String(), kkk.Int()+1)
						cardData = SetUserData("userCard.json", skillIdKey, vvv.Int())
						return true
					})

					card := gjson.Parse(cardData).Get("user_card_by_card_id." + key.String())
					cardList = append(cardList, card.Get("card_master_id").Int())
					cardList = append(cardList, card.Value())

					index++
				}
				return true
			})
		}
		return true
	})

	signBody := GetData("skillEditResult.json")
	signBody, _ = sjson.Set(signBody, "user_model.user_status", GetUserStatus())
	signBody, _ = sjson.Set(signBody, "user_model.user_card_by_card_id", cardList)
	resp := SignResp(ctx.GetString("ep"), signBody, config.SessionKey)
	// fmt.Println(resp)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func SaveDeckLesson(ctx *gin.Context) {
	reqBody := ctx.GetString("reqBody")
	// fmt.Println(reqBody)

	req := gjson.Parse(reqBody).Array()[0]
	deckId := req.Get("deck_id").Int()
	lessonDeck := GetUserData("lessonDeck.json")

	var deckInfo string
	var deckIndex string
	gjson.Parse(lessonDeck).Get("user_lesson_deck_by_id").ForEach(func(key, value gjson.Result) bool {
		if value.IsObject() && value.Get("user_lesson_deck_id").Int() == deckId {
			deckInfo = value.String()
			deckIndex = key.String()
			// fmt.Println("Lesson Deck:", deckInfo)
			return false
		}
		return true
	})

	cardList := req.Get("card_master_ids")
	cardList.ForEach(func(key, value gjson.Result) bool {
		if key.Int()%2 == 0 {
			position := value.String()
			// fmt.Println("Position:", position)

			cardMasterId := cardList.Get(fmt.Sprintf("%d", key.Int()+1)).Int()
			// fmt.Println("Card:", cardMasterId)

			deckInfo, _ = sjson.Set(deckInfo, "card_master_id_"+position, cardMasterId)
			// fmt.Println("New Lesson Deck:", deckInfo)

			SetUserData("lessonDeck.json", "user_lesson_deck_by_id."+deckIndex, gjson.Parse(deckInfo).Value())
			// lessonDeck, _ = sjson.Set(lessonDeck, "user_lesson_deck_by_id."+deckIndex, gjson.Parse(deckInfo).Value())
		}
		return true
	})

	signBody := GetData("saveDeckLesson.json")
	signBody, _ = sjson.Set(signBody, "user_model.user_status", GetUserStatus())
	signBody, _ = sjson.Set(signBody, "user_model.user_lesson_deck_by_id.0", deckId)
	signBody, _ = sjson.Set(signBody, "user_model.user_lesson_deck_by_id.1", gjson.Parse(deckInfo).Value())
	resp := SignResp(ctx.GetString("ep"), signBody, config.SessionKey)
	// fmt.Println(resp)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}
