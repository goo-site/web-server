package main

import (
	"WebServer/internal/utils/log"
)

func main() {
	log.Init()
	log.Info("WebServer start!")
	log.Info("xxx")
	defer func() {
		log.Info("WebServer stop!\n")
	}()
	//router.Run()
}
