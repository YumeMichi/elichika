package db

var (
	DB *Instance
)

func InitDB() {
	DB = GetInstance()
}
