package db

import "xorm.io/xorm"

var (
	DB *Instance

	MainDb  = "assets/main.db"
	MainEng *xorm.Engine
)

func InitDB() {
	DB = GetInstance()

	eng, err := xorm.NewEngine("sqlite", MainDb)
	if err != nil {
		panic(err)
	}
	err = eng.Ping()
	if err != nil {
		panic(err)
	}
	MainEng = eng
	MainEng.SetMaxOpenConns(50)
	MainEng.SetMaxIdleConns(10)
}
