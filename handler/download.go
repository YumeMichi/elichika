package handler

import (
	"elichika/config"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

func GetPackUrl(ctx *gin.Context) {
	reqBody := ctx.GetString("reqBody")
	// fmt.Println(reqBody)

	var packNames []string
	gjson.Parse(reqBody).ForEach(func(key, value gjson.Result) bool {
		if value.IsObject() {
			if err := json.Unmarshal([]byte(value.Get("pack_names").String()), &packNames); err != nil {
				panic(err)
			}
			return false
		}
		return true
	})

	var packUrls []string
	for _, pack := range packNames {
		packUrls = append(packUrls, config.Conf.Settings.CdnServer+"/"+config.Conf.Settings.MasterVersion+"/"+pack)
	}

	packBody, _ := sjson.Set("{}", "url_list", packUrls)
	resp := SignResp(ctx.GetString("ep"), packBody, config.SessionKey)
	// fmt.Println("Response:", resp)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}
