package handler

import (
	"elichika/config"
	"elichika/encrypt"
	"elichika/utils"
	"encoding/json"
	"fmt"
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
