package config

import (
	"elichika/utils"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	_ "modernc.org/sqlite"
)

var (
	// Taken from:
	// https://github.com/RayFirefist/SukuStar_Datamine/blob/master/lib/sifas_api/sifas.py#L120
	// https://github.com/RayFirefist/SukuStar_Datamine/blob/master/lib/sifas_api/sifas.py#L400
	ServerEventReceiverKey = "31f1f9dc7ac4392d1de26acf99d970e425b63335b461e720c73d6914020d6014"
	JaKey                  = "78d53d9e645a0305602174e06b98d81f638eaf4a84db19c756866fddac360c96"

	SessionKey = "12345678123456781234567812345678"

	Conf = &AppConfigs{}
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

func InitConf() {
	Conf = Load("./config.json")

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
		fmt.Println("Configuration file has been generated. Please modify and re-run the program.")
		os.Exit(0)
	}
	c := AppConfigs{}
	err := json.Unmarshal([]byte(utils.ReadAllText(p)), &c)
	if err != nil {
		_ = os.Rename(p, p+".backup"+strconv.FormatInt(time.Now().Unix(), 10))
		_ = DefaultConfigs().Save(p)
		fmt.Println("Configuration file has been generated. Please modify and re-run the program.")
		os.Exit(0)
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
