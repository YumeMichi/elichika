package handler

import (
	"elichika/config"
	"elichika/encrypt"
	"elichika/utils"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/tidwall/sjson"
	"xorm.io/xorm"
)

var (
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
	signBody := fmt.Sprintf("%d,\"%s\",0,%s", time.Now().UnixMilli(), config.MasterVersion, body)
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
	return r
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

func SetUserData(fileName, key string, value any) string {
	userData, err := sjson.Set(GetUserData(fileName), key, value)
	CheckErr(err)

	utils.WriteAllText(userDataPath+fileName, userData)

	return userData
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
