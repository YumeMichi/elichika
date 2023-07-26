package handler

import (
	"elichika/config"
	"elichika/encrypt"
	"elichika/utils"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/tidwall/sjson"
	"xorm.io/xorm"
)

var (
	IsGlobal      = false
	MasterVersion = "b66ec2295e9a00aa"
	StartUpKey    = "5f7IZY1QrAX0D49g"

	MainEng *xorm.Engine

	presetDataPath = "assets/preset/"
	userDataPath   = "assets/userdata/"
)

func init() {
	MainEng = config.MainEng

	os.Mkdir(userDataPath, 0755)
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func SignResp(ep, body, key string) (resp string) {
	signBody := fmt.Sprintf("%d,\"%s\",0,%s", time.Now().UnixMilli(), MasterVersion, body)
	sign := encrypt.HMAC_SHA1_Encrypt([]byte(ep+" "+signBody), []byte(key))
	// fmt.Println(sign)

	resp = fmt.Sprintf("[%s,\"%s\"]", signBody, sign)
	return
}

func GetUserStatus() map[string]any {
	userData := GetUserData("userStatus.json")
	var r map[string]any
	if err := json.Unmarshal([]byte(userData), &r); err != nil {
		panic(err)
	}
	if IsGlobal {
		r["gdpr_version"] = 4
	}
	return r
}

func GetData(fileName string) string {
	presetDataFile := presetDataPath + fileName
	if !utils.PathExists(presetDataFile) {
		panic("File not exists")
	}

	return utils.ReadAllText(presetDataFile)
}

func GetUserData(fileName string) string {
	userDataFile := userDataPath + fileName
	if utils.PathExists(userDataFile) {
		return utils.ReadAllText(userDataFile)
	}

	presetDataFile := presetDataPath + fileName
	if !utils.PathExists(presetDataFile) {
		panic("File not exists")
	}

	userData := utils.ReadAllText(presetDataFile)
	utils.WriteAllText(userDataFile, userData)

	return userData
}

func GetLiveDeckData() string {
	if IsGlobal {
		return GetUserData("liveDeck_gl.json")
	}
	return GetUserData("liveDeck.json")
}

func GetUserAccessoryData() string {
	if IsGlobal {
		return GetData("userAccessory_gl.json")
	}
	return GetData("userAccessory.json")
}

func SetUserData(fileName, key string, value any) string {
	userData, err := sjson.Set(GetUserData(fileName), key, value)
	CheckErr(err)

	utils.WriteAllText(userDataPath+fileName, userData)

	return userData
}

func SetLiveDeckData(key string, value any) string {
	if IsGlobal {
		return SetUserData("liveDeck_gl.json", key, value)
	}
	return SetUserData("liveDeck.json", key, value)
}

func GetPartyInfoByRoleIds(roleIds []int) (partyIcon int, partyName string) {
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
	return
}

func GetRealPartyName(partyName string) (realPartyName string) {
	_, err := MainEng.Table("m_dictionary").Where("id = ?", strings.ReplaceAll(partyName, "k.", "")).
		Cols("message").Get(&realPartyName)
	CheckErr(err)
	return
}

func GetMemberDefaultSuit(cardMasterId int) int {
	var memberMasterId int
	_, err := MainEng.Table("m_card").Where("id = ?", cardMasterId).
		Cols("member_m_id").Get(&memberMasterId)
	CheckErr(err)

	suitMasterId, err := strconv.Atoi(fmt.Sprintf("10%03d1001", memberMasterId))
	CheckErr(err)

	return suitMasterId
}
