package handler

import (
	"elichika/config"
	"elichika/encrypt"
	"elichika/utils"
	"encoding/json"
	"fmt"
	"time"

	"xorm.io/xorm"
)

var (
	MainEng *xorm.Engine
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
	var r map[string]any
	if err := json.Unmarshal([]byte(utils.ReadAllText("assets/userStatus.json")), &r); err != nil {
		panic(err)
	}
	return r
}
