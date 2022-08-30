package main

import (
	"WebServer/internal/config"
	"WebServer/internal/utils/log"
)

func main() {
	config.Init()
	log.Init()
	log.Info("WebServer start!")
	log.Info("xxx")
	defer func() {
		log.Info("WebServer stop!\n")
	}()
	//router.Run()
}
