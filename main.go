package main

import (
	"elichika/config"
	"elichika/db"
	"elichika/patcher"
	"elichika/router"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

func main() {
	// config
	config.InitConf()

	// patcher
	flag.Parse()
	patcher.ApkPatcher()

	// database
	db.InitDB()
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-signalChan
		log.Println("Shutting down...")
		db.DB.Close()
		os.Exit(0)
	}()

	// router
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	router.Router(r)

	go func() {
		if err := r.Run(":" + config.Conf.Settings.ListenPort); err != nil {
			panic(err)
		}
	}()

	fmt.Println("=====================================")
	fmt.Println("[optional] cmdline parameters:")
	fmt.Println("  -apk <file_apk>: Patch apk file")
	fmt.Println("=====================================")
	log.Println("App started")

	select {}
}
