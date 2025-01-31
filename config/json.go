package config

import (
	"elichika/utils"
	"encoding/json"
	"os"
	"strconv"
	"time"
)

type AppConfigs struct {
	AppName  string    `json:"app_name"`
	Settings Settings  `json:"settings"`
	Patcher  []Patcher `json:"patcher"`
}

type Settings struct {
	ListenPort string `json:"listen_port"`
	CdnServer  string `json:"cdn_server"`
}

type LevelDbConfigs struct {
	DataPath string `json:"data_path"`
}

type Patcher struct {
	Target      string `json:"target"`
	Replacement string `json:"replacement"`
}

func DefaultConfigs() *AppConfigs {
	return &AppConfigs{
		AppName: "elichika",
		Settings: Settings{
			ListenPort: "8080",
			CdnServer:  "http://192.168.1.123/static",
		},
		Patcher: []Patcher{
			{
				Target:      "http://127.0.0.1:8080",
				Replacement: "http://192.168.1.123",
			},
			{
				Target:      "http://localhost:8080",
				Replacement: "http://192.168.1.123",
			},
		},
	}
}

func Load(p string) *AppConfigs {
	if !utils.PathExists(p) {
		_ = DefaultConfigs().Save(p)
	}
	c := AppConfigs{}
	err := json.Unmarshal([]byte(utils.ReadAllText(p)), &c)
	if err != nil {
		_ = os.Rename(p, p+".backup"+strconv.FormatInt(time.Now().Unix(), 10))
		_ = DefaultConfigs().Save(p)
	}
	c = AppConfigs{}
	_ = json.Unmarshal([]byte(utils.ReadAllText(p)), &c)
	return &c
}

func (c *AppConfigs) Save(p string) error {
	data, err := json.MarshalIndent(c, "", "	")
	if err != nil {
		return err
	}
	utils.WriteAllText(p, string(data))
	return nil
}
